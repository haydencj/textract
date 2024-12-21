package internal

import (
	"bufio"
	"bytes"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/go-vgo/robotgo"
)

// image processing

func ReadImage(state *State) {
	width := int32(math.Abs(state.activeLoc.X - state.initLoc.X))
	height := int32(math.Abs(state.activeLoc.Y - state.initLoc.Y))

	//TODO: #7 Fix image capture offset
	// use robot go to capture screen
	imageBitmap := robotgo.CaptureScreen(int(state.initLoc.X), int(state.initLoc.Y), int(width), int(height))
	defer robotgo.FreeBitmap(imageBitmap)

	// save image?
	image := robotgo.ToImage(imageBitmap)

	// encode to jpeg
	var imageBuf bytes.Buffer
	err := jpeg.Encode(&imageBuf, image, nil)
	if err != nil {
		log.Fatalln("Image encoding failed:", err)
	}

	// write to file
	file, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	fw := bufio.NewWriter(file)

	fw.Write(imageBuf.Bytes())

}
