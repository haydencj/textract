package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	"log"
	. "screen2text/internal"

	"github.com/getlantern/systray"
	"github.com/go-gl/glfw/v3.3/glfw"
	hook "github.com/robotn/gohook"
)

func main() {
	onExit := func() {
		glfw.Terminate()
	}

	systray.Register(onReady, onExit)

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	// initialize state, window, gl backend and callbacks
	appState := State{Sx: 1, Sy: 1}
	win, backend, cv := Init()
	SetUpCallbacks(&appState, win)

	// initialize clipboard packpage
	InitClipboard()

	var activateSelection bool = false

	// global keyboard event listener
	fmt.Println("--- Please press ctrl + shift + c to start select ---")
	hook.Register(hook.KeyDown, []string{"c", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-c")
		activateSelection = !activateSelection
		hook.End()
	})

	// start hook
	hook.Start()
	// end hook
	defer hook.End()

	// runs every frame
	for !win.ShouldClose() {
		win.MakeContextCurrent()

		// Only run selection logic when active
		if activateSelection {
			// TODO: #2 Move to window size and scaling logic to renderer.go.
			ww, wh := win.GetSize()
			fbw, fbh := win.GetFramebufferSize()
			appState.Sx = float64(fbw) / float64(ww)
			appState.Sy = float64(fbh) / float64(wh)

			glfw.PollEvents()

			// set canvas size
			backend.SetBounds(0, 0, fbw, fbh)

			// call the run function to do all the drawing
			Run(cv, float64(fbw), float64(fbh), &appState)

			// swap back and front buffer
			win.SwapBuffers()
		} else {
			glfw.PollEvents()
		}
	}
}

func onReady() {
	// Set up the tray icon
	//systray.SetIcon(getIcon()) // Replace this with your custom icon if you have one
	systray.SetTitle("Textract")
	systray.SetTooltip("Text Extraction Tool")

	// Add menu items
	startSelection := systray.AddMenuItem("Start Selection", "Activate the selection overlay")
	quit := systray.AddMenuItem("Quit", "Quit the application")

	// Handle menu item clicks
	go func() {
		for {
			select {
			case <-startSelection.ClickedCh:
				fmt.Println("Start Selection clicked!")
				// Call your selection overlay function here
			case <-quit.ClickedCh:
				fmt.Println("Quit clicked!")
				systray.Quit()
			}
		}
	}()
}
