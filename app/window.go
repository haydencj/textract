package app

import (
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/goglbackend"
	"golang.design/x/mainthread"
)

func NewWindow() (*Win, error) {
	var (
		w   = &Win{}
		err error
	)
	// get monitor info
	monitor := glfw.GetPrimaryMonitor()
	vidMode := monitor.GetVideoMode()

	mainthread.Call(func() { setWindowHints(vidMode) })

	mainthread.Call(func() {
		w.win, err = glfw.CreateWindow(vidMode.Width, vidMode.Height,
			"screen2text", nil, nil)
		if err != nil { // window creation failed
			log.Fatalln("failed to create window:", err)
		}
	})
	if err != nil {
		return nil, err
	}

	mainthread.Call(func() { w.win.SetPos(0, 0) })
	w.win.MakeContextCurrent() // changing openGL's state -> changing current context state. one context per thread

	w.SetUpCallbacks()

	return w, nil
}

// Initialize OpenGL Go bindings and configure OpenGL settings
func SetUpGL() (*goglbackend.GoGLBackend, *canvas.Canvas) {
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

	cv := canvas.New(backend)
	return backend, cv
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
