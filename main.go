package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	"log"
	. "screen2text/app"

	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.design/x/mainthread"
)

// TRY IT WITHOUT MAINTHREAD PACKAGE!
// YOU KNOW YOU CAN GET SYSTRAY FORK AND GLFW WORKING TOGETHER WITHOUT MAINTHREAD PACKAGE.
// JUST CREATE WINDOW AT START WITH SYSTRAY CREATION AND HIDE THE WINDOW.
// START THE SELECTION OVERLAY DRAWING WHEN HOTKEY IS PRESSED

func main() { mainthread.Init(fn) }

func fn() {
	InitClipboard()

	// Start systray
	fmt.Println("Starting systray...")
	startTray, endTray := StartSystray()
	defer endTray()

	// Initialize GLFW once
	var err error
	mainthread.Call(func() { err = glfw.Init() })
	if err != nil {
		fmt.Println("failed to initialize GLFW")
		return
	}
	defer mainthread.Call(func() { glfw.Terminate() })

	var win *Win
	// Schedule the window creation on the main thread:
	mainthread.Call(func() {
		win, err = NewWindow()
	})
	if err != nil {
		log.Fatalln("Failed to create OpenGL window:", err)
	}

	// Create hotkey channel
	hotkeyChan := RegisterHotkey()

	//Run systray on main thread
	mainthread.Call(func() {
		startTray()
	})

	for range hotkeyChan {
		fmt.Println("Hotkey pressed, starting selection...")
		win.StartSelection()
	}

}

// func fn() {
// 	InitClipboard()

// 	// Initialize GLFW once at startup
// 	var err error
// 	mainthread.Call(func() { err = glfw.Init() })
// 	if err != nil {
// 		fmt.Println("failed to initialize GLFW")
// 		return
// 	}
// 	defer mainthread.Call(func() { glfw.Terminate() })
// 	fmt.Println("GLFW initialized successfully")

// 	// Main application loop
// 	for {
// 		// Start systray and wait for hotkey
// 		systrayChan := make(chan struct{})
// 		go runSystray(systrayChan)
// 		<-systrayChan // Wait for systray to exit

// 		// Run selection overlay
// 		StartSelection()
// 	}
// }

func runSystray(done chan struct{}) {
	hotkeyChan := RegisterHotkey()
	startTray, endTray := StartSystray()
	defer endTray()

	mainthread.Call(func() { startTray() })

	// Wait for hotkey
	<-hotkeyChan
	fmt.Println("Hotkey pressed!")
	done <- struct{}{} // Signal main loop that systray is done
}

// func fn() {
// 	InitClipboard()
// 	hotkeyChan := RegisterHotkey()

// 	startTray, endTray := StartSystray()
// 	defer endTray()

// 	mainthread.Call(func() { startTray() })

// 	go func() {
// 		for {
// 			select {
// 			case <-hotkeyChan:
// 				fmt.Println("Hotkey pressed!")
// 				systray.Quit()
// 			}
// 		}
// 	}()

// 	var err error

// 	mainthread.Call(func() { err = glfw.Init() })
// 	if err != nil {
// 		fmt.Println("failed to initialize GLFW")
// 	}
// 	fmt.Println("GLFW initialized successfully")

// 	StartSelection()

// 	defer mainthread.Call(func() { glfw.Terminate() })

// 	select {}
// }
