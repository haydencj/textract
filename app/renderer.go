package app

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tfriedel6/canvas"
)

// Run runs the given window and blocks until it is destroyed.
func (w *Win) Run() {
	w.win.Show()
	for !w.win.ShouldClose() {

		w.win.MakeContextCurrent()

		// Only run selection logic when active
		// TODO: #2 Move to window size and scaling logic to renderer.go.
		ww, wh := w.win.GetSize()
		fbw, fbh := w.win.GetFramebufferSize()
		w.state.Sx = float64(fbw) / float64(ww)
		w.state.Sy = float64(fbh) / float64(wh)

		// set canvas size
		w.backend.SetBounds(0, 0, fbw, fbh)

		// call the run function to do all the drawing
		Draw(w.cv, float64(fbw), float64(fbh), w.state)

		// swap back and front buffer
		w.win.SwapBuffers()

		glfw.PollEvents()
	}
	// This function must be called from the mainthread.
	// mainthread.Call(w.win.Destroy)
	w.win.Destroy()
}

// TODO: #5 Consider breaking Run into separate drawing functions.
func Draw(cv *canvas.Canvas, w, h float64, s *State) {
	dynamicHeight := s.GLMouse.activeLoc.Y - s.GLMouse.initLoc.Y
	dynamicWidth := s.GLMouse.activeLoc.X - s.GLMouse.initLoc.X

	cv.ClearRect(0, 0, w, h) // 'refreshes' canvas, clears back buffer

	// semi-transparent background
	cv.SetFillStyle(20, 22, 22, 0.7)
	cv.FillRect(0, 0, w, h)

	// punch out (selection area)
	cv.ClearRect(s.GLMouse.initLoc.X, s.GLMouse.initLoc.Y, dynamicWidth, dynamicHeight)
	// 'border' for selection area
	cv.SetStrokeStyle(255, 255, 255)
	cv.StrokeRect(s.GLMouse.initLoc.X, s.GLMouse.initLoc.Y, dynamicWidth, dynamicHeight)

}
