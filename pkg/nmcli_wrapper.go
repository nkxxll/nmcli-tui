package main

import (
	"os/exec"
	"strings"
)

/* lists all available wifis */
func wifi_list() []string {
    var list []string
    cmd := exec.Command("nmcli", "dev", "wifi", "list")
    stdout, err := cmd.Output()

    if err != nil {
        panic("Command nmcli dev wifi list failed")
    }

    list = strings.Split(string(stdout), "\n")
    return list
}

/** 
connects to a network via the nmcli tool
args: passwd, ssid
returns: err
*/
func wifi_connect(passwd string, ssid string) error {
    // sudo nmcli dev wifi connect network-ssid
    ssid = "\"" + ssid + "\""
    passwd = "\"" + passwd + "\""
    cmd := exec.Command("sudo", "nmcli", "dev", "wifi", "connect", ssid, "password", passwd)
    // TODO: probably need the output have to wait and see
    err := cmd.Run()
    return err
}

/** 
show saved connections
args: nil
returns: string output
*/
func wifi_show_saved() string {
    var res string
    cmd := exec.Command("nmcli", "con", "show")
    // TODO: probably need the output have to wait and see
    stdout, err := cmd.Output()
    if err != nil {
        panic("There was an error while showing saved wifis")
    }
    res = string(stdout)
    return res
}

/** 
kills connection to a known ssid
args: ssid
returns: err
*/
func wifi_con_down(ssid string) error {
    cmd := exec.Command("nmcli", "con", "down", ssid)
    // TODO: probably need the output have to wait and see
    err := cmd.Run()
    return err
}
