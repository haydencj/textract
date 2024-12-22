package internal

import (
	"fmt"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-vgo/robotgo"
)

// TODO: #3 Remove global variables. Decide if should put in state or else where.
var (
	isMouseHeld    bool
	isMouseRelease bool
)

func SetUpCallbacks(state *State, window *glfw.Window) {

	// define behavior when mouse is pressed and released
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		mouseButtonCallback(w, button, action, state)
	})

	// when cursor moves - update active location
	window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
		cursorPosCallback(xpos, ypos, state)
	})

	// user can close window via escape key
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		keyCallback(w, key, action)
	})
}

// TODO: #4 Clean up callback functions. Consider making them method receivers on state.
// TODO: #8 Compare robotgo mouse position w/ openGL mouse position.
func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, s *State) {
	if button == glfw.MouseButton1 {
		isMouseHeld = (action == glfw.Press)
		isMouseRelease = (action == glfw.Release)

		// on click - initialize locations
		if isMouseHeld {
			glX, glY := w.GetCursorPos()
			sysX, sysY := robotgo.Location()

			newGLCoord := Coordinate[float64]{glX, glY}
			newSysCoord := Coordinate[int]{sysX, sysY}

			s.GLMouse.setLocation(&s.GLMouse.initLoc, newGLCoord)
			s.GLMouse.initLoc.scale(s)
			s.GLMouse.setLocation(&s.GLMouse.activeLoc, newGLCoord)
			s.GLMouse.activeLoc.scale(s)

			s.SystemMouse.setLocation(&s.SystemMouse.initLoc, newSysCoord)
		}

		// on release - finalize active location (initial doesn't change)
		if isMouseRelease {
			glX, glY := w.GetCursorPos()
			sysX, sysY := robotgo.Location()

			newGLCoord := Coordinate[float64]{glX, glY}
			newSysCoord := Coordinate[int]{sysX, sysY}

			s.GLMouse.setLocation(&s.GLMouse.activeLoc, newGLCoord)
			s.GLMouse.activeLoc.scale(s)

			s.SystemMouse.setLocation(&s.SystemMouse.activeLoc, newSysCoord)

			// print width, height of rectangle
			fmt.Println(math.Abs(s.GLMouse.activeLoc.X-s.GLMouse.initLoc.X), math.Abs(s.GLMouse.activeLoc.Y-s.GLMouse.initLoc.Y))

			ReadImage(s)

		}
	}
}

func cursorPosCallback(xpos, ypos float64, s *State) {
	if isMouseHeld {
		s.GLMouse.activeLoc.X, s.GLMouse.activeLoc.Y = xpos, ypos
		s.GLMouse.activeLoc.scale(s)
	}
}

func keyCallback(w *glfw.Window, key glfw.Key, action glfw.Action) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

// getLocationAndScale sets a coordinate location.
// The coord parameter should be a pointer to either initLoc or activeLoc.
func (m *Mouse[T]) setLocation(coord *Coordinate[T], newCoord Coordinate[T]) {
	coord.X = newCoord.X
	coord.Y = newCoord.Y
}
