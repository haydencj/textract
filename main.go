package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	"log"
	. "screen2text/app"

	_ "embed" // Import embed for embedding the icon

	"github.com/getlantern/systray"
	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.design/x/mainthread"
)

// Embed the tray icon
//
//go:embed assets/icon.png
var iconData []byte

// // Channels for inter-thread communication
// var showWindowChan = make(chan struct{})
// var quitChan = make(chan struct{})

// // Window state management
// var windowVisible bool

func Init() (err error) {
	mainthread.Call(func() { err = glfw.Init() })
	//mainthread.Call(func() { systray.Run(onReady, onExit) })
	return
}

func main() {
	mainthread.Init(fn)
}

func fn() {

	err := Init()
	if err != nil {
		panic(err)
	}
	defer Terminate()

	window, err := NewWindow()
	if err != nil {
		panic(err)
	}

	InitClipboard()

	// Start systray
	//mainthread.Call(func() { systray.Register(onReady, nil) })

	// var activateSelection bool = false

	// // global keyboard event listener
	// fmt.Println("--- Please press ctrl + shift + c to start select ---")
	// hook.Register(hook.KeyDown, []string{"c", "ctrl", "shift"}, func(e hook.Event) {
	// 	fmt.Println("ctrl-shift-c")
	// 	activateSelection = !activateSelection
	// 	hook.End()
	// })

	// // start hook
	// hook.Start()
	// // end hook
	// defer hook.End()

	// runs every frame
	window.Run()

}

func onReady() {
	log.Println("Systray is ready")
	// Set up the tray icon
	//systray.SetIcon(iconData)
	systray.SetTitle("Textract")
	systray.SetTooltip("Text Extraction Tool")

	startSelection := systray.AddMenuItem("Start Selection", "Activate the selection overlay")
	quit := systray.AddMenuItem("Quit", "Quit the application")

	// Handle menu items in a separate goroutine
	go func() {
		for {
			select {
			case <-startSelection.ClickedCh:
				fmt.Println("Start Selection clicked!")
				//showWindowChan <- struct{}{} // Signal to show window
			case <-quit.ClickedCh:
				fmt.Println("Quit clicked!")
				//quitChan <- struct{}{} // Signal to quit
				return
			}
		}
	}()
}

func onExit() {
	// Perform cleanup if needed
	fmt.Println("Exiting application...")
	Terminate()
}

func Terminate() {
	fmt.Println("Terminating...")
	mainthread.Call(glfw.Terminate)
}
