package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: not so nice should be relative
const listHeight = 14

// Here when adding style to the view maybe Create a struct instead of Variables ?ex:
/**
type Styles struct{
  titleStyle        = lipgloss.Style
	itemStyle         = lipgloss.Style
	selectedItemStyle =  lipgloss.Style
	paginationStyle   = lipgloss.Style
	helpStyle         = lipgloss.Style
	quitTextStyle     = lipgloss.Style
}

// than here we use the style struct to create a Default Style
func DefaultStyles() *Styles{
  s:= new(Styles)
  s.titleStyle = lipgloss.NewStyle().MarginLeft(2)
}

*/
var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("5"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			// TODO: probably change to something nice
			return selectedItemStyle.Render("-> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	options  list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.options.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.options.SelectedItem().(Item)
			if ok {
				m.choice = string(i)
			}
			// TODO: here comes the nmcli logic
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)
	return m, cmd
}

func (m model) View() string {
	// Of course here instead of matching on the Strings like here case Network ... we should match on i think command
	// I guess something like command.listNetworks , command.connectToNetwork and so on ?
	if m.choice != "" {
		switch m.choice {
		case "List Networks":
			// here we should not quit but render a new view with a list of the search results of the nmcli cmd
			return quitTextStyle.Render(fmt.Sprint("Instead of this the programm should execute in the background the nmcli command and return a new view with the Network Connections that are possible !"))
		case "Connect to a Network":
			return quitTextStyle.Render(fmt.Sprintf("With this %s I am not sure what to do ?:)", m.choice))
		case "Quit":
			return quitTextStyle.Render(fmt.Sprintf("%s ... ? Have A nice day:)", m.choice))
		default:
			return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
		}
	}
	if m.quitting {
		return quitTextStyle.Render("Quiting NMCLI !")
	}
	return "\n" + m.options.View()
}

func main() {
	options := []list.Item{
		Item("List Networks"),
		Item("Connect to a Network"),
		Item("Quit"),
	}
	const defaultWidth = 20

	l := list.New(options, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "NMCLI WRAPPER"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{options: l}

	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
