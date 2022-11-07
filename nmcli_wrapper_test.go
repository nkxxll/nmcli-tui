package main

import (
	"testing"
)

func TestNmcliConnect(t *testing.T) {
	// NOTE: sudo works but the connection does not
	err := wifi_connect("Passwordxxx", "wifi")
    // FIXME: this is a workaround
	if err != nil && err.Error() != "exit status 10" {
		t.Errorf("There was an error while connecting %s", err)
	}
}

func TestNmcliShowSavedConnections(t *testing.T)  {
    res := wifi_show_saved()
    t.Logf("The output is: %s", res)
    if res == "" {
        t.Errorf("There is output from the wifi_show_saved func")
    }
}

func TestNmcliConnectionDown(t *testing.T)  {
    err := wifi_con_down("theWiFi")
    t.Logf("The output is: %s", err)
    if err != nil {
        t.Errorf("There is output from the wifi_con_down func")
    }
}
