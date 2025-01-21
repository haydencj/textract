package app

import (
	"fmt"

	"github.com/energye/systray"
	"golang.design/x/mainthread"
)

func (win *Win) StartSelection() {
	fmt.Println("Running OpenGL Selection...")

	// Run the OpenGL selection
	mainthread.Call(func() { win.Run() })

	fmt.Println("OpenGL Selection complete.")
}

func StartSystray() (func(), func()) {
	var (
		startTray func()
		endTray   func()
	)

	fmt.Println("About to run systray...")

	mainthread.Call(func() { startTray, endTray = systray.RunWithExternalLoop(OnReady, OnExit) })

	return startTray, endTray
}
