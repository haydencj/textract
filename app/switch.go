package app

import (
	"fmt"
	"log"

	"github.com/energye/systray"
	"golang.design/x/mainthread"
)

func StartSelection() {
	var (
		win *Win
		err error
	)

	fmt.Println("Running OpenGL Selection...")

	// Schedule the window creation on the main thread:
	mainthread.Call(func() {
		win, err = NewWindow()
	})

	if err != nil {
		log.Fatalln("Failed to create OpenGL window:", err)
	}

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
