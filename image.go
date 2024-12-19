package main

import (
	"fmt"
	"math"

	"github.com/go-gl/gl/v2.1/gl"
)

// image processing

func ReadImage(initLoc Coordinate, activeLoc Coordinate) {
	width := int32(math.Abs(activeLoc.X - initLoc.X))
	height := int32(math.Abs(activeLoc.Y - initLoc.Y))

	pixels := make([]byte, width*height*4) // 4 bytes for each pixel (RGBA)
	gl.ReadPixels(int32(initLoc.X), int32(initLoc.Y), width, height, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	fmt.Println(pixels)
}
