package main

// TODO: Remove wildcard import for internal package.
import (
	"log"
	. "screen2text/internal"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// initialize state, window, gl backend and callbacks
	appState := State{Sx: 1, Sy: 1}
	win, backend, cv := Init()
	SetUpCallbacks(&appState, win)

	// runs every frame
	for !win.ShouldClose() {
		win.MakeContextCurrent()

		// TODO: Move to window size and scaling logic to renderer.go.
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
	}
}
