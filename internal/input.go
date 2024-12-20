package internal

import (
	"fmt"
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// TODO: Remove global variables. Decide if should put in state or else where.
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

// TODO: Clean up callback functions. Consider making them method receivers on state.
func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, state *State) {
	if button == glfw.MouseButton1 {
		isMouseHeld = (action == glfw.Press)
		isMouseRelease = (action == glfw.Release)

		// on click - initialize locations
		if isMouseHeld {
			state.initLoc.X, state.initLoc.Y = w.GetCursorPos()
			state.initLoc.scale(state)
			state.activeLoc.X, state.activeLoc.Y = w.GetCursorPos()
			state.activeLoc.scale(state)

		}

		// on release - finalize active location (initial doesn't change)
		if isMouseRelease {
			state.activeLoc.X, state.activeLoc.Y = w.GetCursorPos()
			state.activeLoc.scale(state)

			// print width, height of rectangle
			fmt.Println(math.Abs(state.activeLoc.X-state.initLoc.X), math.Abs(state.activeLoc.Y-state.initLoc.Y))

		}
	}
}

func cursorPosCallback(xpos, ypos float64, state *State) {
	if isMouseHeld {
		state.activeLoc.X, state.activeLoc.Y = xpos, ypos
		state.activeLoc.scale(state)
	}
}

func keyCallback(w *glfw.Window, key glfw.Key, action glfw.Action) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}
