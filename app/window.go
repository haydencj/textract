package app

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

func NewWindow() (*Win, error) {
	var (
		w   = &Win{}
		err error
	)

	fmt.Println("Begin window creation.")

	// initialize state
	w.state = &State{Sx: 1, Sy: 1}

	// get monitor info
	monitor := glfw.GetPrimaryMonitor()
	if monitor == nil {
		return nil, fmt.Errorf("failed to get primary monitor")
	}
	fmt.Println("Got primary monitor")

	vidMode := monitor.GetVideoMode()
	if vidMode == nil {
		return nil, fmt.Errorf("failed to get video mode")
	}
	fmt.Printf("Got video mode: %dx%d@%dHz\n", vidMode.Width, vidMode.Height, vidMode.RefreshRate)

	// Set window hints before creation
	setWindowHints(vidMode)
	fmt.Println("Window hints set")

	// Try creating window
	fmt.Printf("Attempting to create window: %dx%d\n", vidMode.Width, vidMode.Height)
	w.win, err = glfw.CreateWindow(vidMode.Width, vidMode.Height,
		"screen2text", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	w.win.Hide()
	w.win.SetPos(0, 0)
	w.win.MakeContextCurrent()

	// initialize OpenGL backend and canvas
	err = w.InitGLBackend()
	if err != nil {
		return nil, err
	}

	// Set up callbacks
	w.SetUpCallbacks()

	fmt.Println("Window created successfully")
	return w, nil
}

// Initialize OpenGL Go bindings and configure OpenGL settings
func (w *Win) InitGLBackend() error {
	err := gl.Init()
	if err != nil {
		return err
	}

	// set vsync on, enable multisample (if available) (OPTIONAL???)
	glfw.SwapInterval(1)
	gl.Enable(gl.MULTISAMPLE)

	// --- BLENDING ---
	gl.Enable(gl.BLEND)
	// the destination is what's already on your "canvas" (the framebuffer), and the source is what you're about to draw
	gl.BlendEquation(gl.FUNC_SUBTRACT) // source - destination

	// load GL backend
	backend, err := goglbackend.New(0, 0, 0, 0, nil)
	if err != nil {
		return err
	}

	cv := canvas.New(backend)

	w.backend = backend
	w.cv = cv

	fmt.Println("GL backend initialized.")
	return nil
}

func setWindowHints(vidMode *glfw.VidMode) {

	// for windows fullscreen
	glfw.WindowHint(glfw.RedBits, vidMode.RedBits)
	glfw.WindowHint(glfw.GreenBits, vidMode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, vidMode.BlueBits)
	glfw.WindowHint(glfw.RefreshRate, vidMode.RefreshRate)

	// the stencil size setting is required for the canvas to work
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)

	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True) // transparent window
	glfw.WindowHint(glfw.Decorated, glfw.False)             // does window info (close button) exist
	glfw.WindowHint(glfw.Floating, glfw.True)               // topmost window
	glfw.WindowHint(glfw.Resizable, glfw.False)             // is window resizable
	glfw.WindowHint(glfw.Maximized, glfw.True)              // maximize window
}
