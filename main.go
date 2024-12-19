package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

type Coordinate struct {
	X, Y float64
}

var (
	isMouseHeld    bool = false
	isMouseRelease bool = false
	initLoc        Coordinate
	activeLoc      Coordinate
	sx             float64 = 1
	sy             float64 = 1
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	defer glfw.Terminate()

	monitor := glfw.GetPrimaryMonitor()
	vidMode := monitor.GetVideoMode()

	// for windowed fullscreen.. ?
	// glfw.WindowHint(glfw.RedBits, vidMode.RedBits)
	// glfw.WindowHint(glfw.GreenBits, vidMode.GreenBits)
	// glfw.WindowHint(glfw.BlueBits, vidMode.BlueBits)
	// glfw.WindowHint(glfw.RefreshRate, vidMode.RefreshRate)

	// the stencil size setting is required for the canvas to work
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)

	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True) // transparent window
	glfw.WindowHint(glfw.Decorated, glfw.False)             // does window info (close button) exist
	glfw.WindowHint(glfw.Floating, glfw.True)               // topmost window
	glfw.WindowHint(glfw.Resizable, glfw.False)             // is window resizable

	window, err := glfw.CreateWindow(vidMode.Width, vidMode.Height, "screen2text", nil, nil)
	if err != nil { // window creation failed
		panic(err)
	}
	window.MakeContextCurrent() // changing openGL's state -> changing current context state. one context per thread

	// init GL
	err = gl.Init()
	if err != nil {
		log.Fatalf("Error initializing GL: %v", err)
	}

	// set vsync on, enable multisample (if available) (OPTIONAL???)
	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	// blending
	gl.Enable(gl.BLEND)
	// the destination is what's already on your "canvas" (the framebuffer), and the source is what you're about to draw
	gl.BlendEquation(gl.FUNC_SUBTRACT) // source - destination

	// load GL backend
	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error loading canvas GL assets: %v", err)
	}

	// callbacks

	// when cursor moves - update active location
	window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
		if isMouseHeld {
			activeLoc.X, activeLoc.Y = xpos, ypos
			activeLoc.scale()
		}

	})

	// define behavior when mouse is pressed and released
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		if button == glfw.MouseButton1 {
			isMouseHeld = (action == glfw.Press)
			isMouseRelease = (action == glfw.Release)

			// on click - initialize locations
			if isMouseHeld {
				initLoc.X, initLoc.Y = window.GetCursorPos()
				initLoc.scale()
				activeLoc.X, activeLoc.Y = window.GetCursorPos()
				activeLoc.scale()

			}

			// on release - finalize active location (initial doesn't change)
			if isMouseRelease {
				activeLoc.X, activeLoc.Y = window.GetCursorPos()
				activeLoc.scale()

			}
		}

	})

	// user can close window via escape key
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape && action == glfw.Press {
			window.SetShouldClose(true)
		}
	})

	// initialize canvas with zero size, since size is set in main loop
	cv := canvas.New(backend)

	for !window.ShouldClose() {
		window.MakeContextCurrent()

		// find window size and scaling
		ww, wh := window.GetSize()
		fbw, fbh := window.GetFramebufferSize()
		sx = float64(fbw) / float64(ww)
		sy = float64(fbh) / float64(wh)

		glfw.PollEvents()

		// set canvas size
		backend.SetBounds(0, 0, fbw, fbh)

		// call the run function to do all the drawing
		run(cv, float64(fbw), float64(fbh))

		// swap back and front buffer
		window.SwapBuffers()
	}
}

func run(cv *canvas.Canvas, w, h float64) {
	cv.ClearRect(0, 0, w, h) // 'refreshes' canvas

	// semi-transparent background
	cv.SetFillStyle(20, 22, 22, 0.7)
	cv.FillRect(0, 0, w, h)

	// punch out (selection area)
	cv.ClearRect(initLoc.X, initLoc.Y, activeLoc.X-initLoc.X, activeLoc.Y-initLoc.Y)
	// 'border' for selection area
	cv.SetStrokeStyle(255, 255, 255)
	cv.StrokeRect(initLoc.X, initLoc.Y, activeLoc.X-initLoc.X, activeLoc.Y-initLoc.Y)

}

func (coord *Coordinate) scale() {
	coord.X = coord.X * sx
	coord.Y = coord.Y * sy
}
