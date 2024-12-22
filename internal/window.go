package internal

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// Initialize GLFW window, OpenGL context, OpenGL backend, and Canvas
func Init() (*glfw.Window, *goglbackend.GoGLBackend, *canvas.Canvas) {
	// --- Initialize Window and OpenGL Context ---

	// get monitor info
	monitor := glfw.GetPrimaryMonitor()
	vidMode := monitor.GetVideoMode()

	// the stencil size setting is required for the canvas to work
	glfw.WindowHint(glfw.StencilBits, 8)
	glfw.WindowHint(glfw.DepthBits, 0)

	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True) // transparent window
	glfw.WindowHint(glfw.Decorated, glfw.False)             // does window info (close button) exist
	glfw.WindowHint(glfw.Floating, glfw.True)               // topmost window
	glfw.WindowHint(glfw.Resizable, glfw.False)             // is window resizable
	glfw.WindowHint(glfw.Maximized, glfw.True)              // maximize window

	window, err := glfw.CreateWindow(vidMode.Width, vidMode.Height, "screen2text", nil, nil)
	if err != nil { // window creation failed
		log.Fatalln("failed to create window:", err)
	}

	window.SetPos(0, 0)
	window.MakeContextCurrent() // changing openGL's state -> changing current context state. one context per thread

	backend := setUpGL()

	// initialize canvas with zero size, since size is set in main loop
	cv := canvas.New(backend)

	return window, backend, cv
}

// Initialize OpenGL Go bindings and configure OpenGL settings
func setUpGL() *goglbackend.GoGLBackend {
	err := gl.Init()
	if err != nil {
		log.Fatalln("failed to initialize GL:", err)
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
		log.Fatalln("error loading canvas GL assets:", err)
	}

	return backend
}
