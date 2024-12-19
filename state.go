package main

// State holds all shared application state.
type State struct {
	InitLoc   Coordinate
	ActiveLoc Coordinate
}

type Coordinate struct {
	X, Y float64
}

// focus on refactoring for state management
