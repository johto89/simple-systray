package main

import (
	"fmt"
	"io/ioutil"

	"github.com/getlantern/systray"
	"github.com/tawesoft/golib/v2/dialog"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetIcon(getIcon("assets/fire.ico"))

	scan := systray.AddMenuItem("scan", "scan")
	check := systray.AddMenuItem("check", "check")
	test := systray.AddMenuItem("test", "test")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			select {
			case <-scan.ClickedCh:
				dialog.Alert("Scan")
			case <-check.ClickedCh:
				dialog.Alert("Check")
			case <-test.ClickedCh:
				dialog.Alert("Test")
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
