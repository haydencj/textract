package app

import (
	_ "embed" // Import embed for embedding the icon
	"fmt"

	"github.com/energye/systray"

	//"github.com/getlantern/systray"
	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.design/x/mainthread"
)

// Embed the tray icon
//
//go:dembed /assets/icon.png
var iconData []byte

func OnReady() {
	fmt.Println("Systray is ready")
	// Set up the tray icon
	//systray.SetIcon(iconData)
	systray.SetTitle("Textract")
	systray.SetTooltip("Text Extraction Tool")

	systray.CreateMenu()

	startSelection := systray.AddMenuItem("Start Selection", "Activate the selection overlay")
	quit := systray.AddMenuItem("Quit", "Quit the application")

	//hotkeyChan := RegisterHotkey()

	startSelection.Click(func() {
		fmt.Println("Start selection clicked")
		// NewWindow()
	})

	quit.Click(func() {
		fmt.Println("Quit clicked")
	})

	// Handle menu items in a separate goroutine (event loop)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-hotkeyChan:
	// 			fmt.Println("Hotkey!")
	// 			systray.Quit()
	// 		case <-startSelection.ClickedCh:
	// 			fmt.Println("Start Selection clicked!")
	// 		case <-quit.ClickedCh:
	// 			fmt.Println("Quit clicked!")
	// 			//quitChan <- struct{}{} // Signal to quit
	// 			return
	// 		}
	// 	}
	// }()
}

func OnExit() {
	// Perform cleanup if needed
	fmt.Println("Exiting systray...")
}

func Terminate() {
	fmt.Println("Terminating...")
	mainthread.Call(glfw.Terminate)
}
