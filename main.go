package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	"log"
	. "screen2text/app"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// TRY IT WITHOUT MAINTHREAD PACKAGE!
// YOU KNOW YOU CAN GET SYSTRAY FORK AND GLFW WORKING TOGETHER WITHOUT MAINTHREAD PACKAGE.
// JUST CREATE WINDOW AT START WITH SYSTRAY CREATION AND HIDE THE WINDOW.
// START THE SELECTION OVERLAY DRAWING WHEN HOTKEY IS PRESSED

// func main() { mainthread.Init(fn) }

func main() {
	InitClipboard()
	// Create hotkey channel
	hotkeyChan := RegisterHotkey()

	// Start systray
	fmt.Println("Starting systray...")
	startTray, endTray := StartSystray()
	defer endTray()

	// Initialize GLFW once
	var err error
	//mainthread.Call(func() { err = glfw.Init() })
	err = glfw.Init()
	if err != nil {
		fmt.Println("failed to initialize GLFW")
		return
	}
	//defer mainthread.Call(func() { glfw.Terminate() })
	defer glfw.Terminate()

	var win *Win

	win, err = NewWindow()
	if err != nil {
		log.Fatalln("Failed to create OpenGL window:", err)
	}

	startTray()

	// Poll events until hotkey is pressed
	running := true
	for running {
		select {
		case <-hotkeyChan:
			running = false
			fmt.Println("Hotkey pressed, moving to selection loop...")
			win.StartSelection()
		default:
			glfw.PollEvents()
		}
	}

}
