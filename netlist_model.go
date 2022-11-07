package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type netlistModel struct {
	list     list.Model
	choice   string
	quitting bool
}

const (
	defaultWidth = 20
)

func NewNetlistModel(netlist []string) tea.Model {
	var m netlistModel
	// TODO: implement
	var items []list.Item
	for i, item := range wifi_list() {
		if i != 0 {
			items = append(items, Item(item))
		}
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "WiFi Options list:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m = netlistModel{list: l}
	return m
}

func (m netlistModel) Init() tea.Cmd {
	return nil
}

func (m netlistModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i)
			}
			// TODO: here comes the nmcli logic
			list := strings.Split(" ", strings.TrimSpace(strings.Trim("*", m.choice)))

			// TODO: passwd with input model
			// err := wifi_connect(passwd, list[2])
			// if err != nil {
			//     panic("There was an error while conntecting")
			// }
			fmt.Println("You chose: %s", list[2])
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m netlistModel) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}
