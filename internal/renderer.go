package internal

import "github.com/tfriedel6/canvas"

// TODO: #5 Consider breaking Run into separate drawing functions.
func Run(cv *canvas.Canvas, w, h float64, s *State) {
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
