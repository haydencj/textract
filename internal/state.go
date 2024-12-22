package internal

// Coordinate represents a point in a coordinate system
type Coordinate[T float64 | int] struct {
	X, Y T
}

// Mouse holds both initial and active coordinates
type Mouse[T float64 | int] struct {
	initLoc   Coordinate[T]
	activeLoc Coordinate[T]
}

// holds all application state
type State struct {
	GLMouse     Mouse[float64]
	SystemMouse Mouse[int]
	//imageBuffer []byte
	Sx float64
	Sy float64
}

// TODO: #6 Remove this method. Move scaling logic to SetInitLoc() and SetActiveLoc() so that coordinates are always scaled properly.
func (coord *Coordinate[float64]) scale(state *State) {
	coord.X *= float64(state.Sx)
	coord.Y *= float64(state.Sy)
}
