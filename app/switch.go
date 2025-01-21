package app

import (
	"fmt"
	"log"

	"github.com/energye/systray"
	//"github.com/getlantern/systray"
)

func SwitchUI() {
	//fmt.Println("Quitting Systray...")

	fmt.Println("Running OpenGL Selection...")
	window, err := NewWindow()
	if err != nil {
		log.Fatalln("Failed to create OpenGL window:", err)
	}

	// Run the OpenGL selection
	window.Run()

	fmt.Println("OpenGL Selection complete. Restarting Systray...")
	//StartSystray()
}

func StartSystray() (func(), func()) {
	var (
		startTray func()
		endTray   func()
	)

	fmt.Println("About to run systray...")

	// mainthread.Call(func() { startTray, endTray = systray.RunWithExternalLoop(OnReady, OnExit) })
	startTray, endTray = systray.RunWithExternalLoop(OnReady, OnExit)

	fmt.Println("mainthread.Call returned!") // <-- And this

	return startTray, endTray
}
