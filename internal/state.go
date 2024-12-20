package internal

type Coordinate struct {
	X, Y float64
}

// holds all application state
type State struct {
	initLoc   Coordinate
	activeLoc Coordinate
	//imageBuffer []byte
	Sx float64
	Sy float64
}

// TODO: Remove this method. Move scaling logic to SetInitLoc() and SetActiveLoc() so that coordinates are always scaled properly.
func (coord *Coordinate) scale(state *State) {
	coord.X *= state.Sx
	coord.Y *= state.Sy
}
