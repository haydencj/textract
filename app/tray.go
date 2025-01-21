package app

import (
	_ "embed" // Import embed for embedding the icon
	"fmt"

	"github.com/energye/systray"
)

// Embed the tray icon
//
//go:dembed /assets/icon.png
var iconData []byte

func OnReady() {
	fmt.Println("Systray onReady is running")
	// Set up the tray icon
	//systray.SetIcon(iconData)
	systray.SetTitle("Textract")
	systray.SetTooltip("Text Extraction Tool")

	systray.CreateMenu()

	startSelection := systray.AddMenuItem("Start Selection", "Activate the selection overlay")
	quit := systray.AddMenuItem("Quit", "Quit the application")

	startSelection.Click(func() {
		fmt.Println("Start selection clicked")
		StartSelection()
	})

	quit.Click(func() {
		fmt.Println("Quit clicked")
	})
}

func OnExit() {
	// Perform cleanup if needed
	fmt.Println("Exiting systray...")
}
