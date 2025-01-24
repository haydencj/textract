package app

import (
	"fmt"

	"github.com/energye/systray"
)

func (win *Win) StartSelection() {
	fmt.Println("Running OpenGL Selection...")

	// Run the OpenGL selection
	win.Run()

	fmt.Println("OpenGL Selection complete.")
}

func StartSystray() (func(), func()) {
	var (
		startTray func()
		endTray   func()
	)

	fmt.Println("About to run systray...")

	startTray, endTray = systray.RunWithExternalLoop(OnReady, OnExit)

	return startTray, endTray
}
