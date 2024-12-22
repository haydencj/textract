package internal

import (
	"bufio"
	"image/png"
	"log"
	"os"

	"github.com/go-vgo/robotgo"
)

// image processing

func ReadImage(s *State) {
	activeLoc := s.SystemMouse.activeLoc
	initLoc := s.SystemMouse.initLoc

	width := abs(activeLoc.X - initLoc.X)
	height := abs(activeLoc.Y - initLoc.Y)

	log.Println("image wh:", width, height)

	//TODO: #7 Fix image capture offset
	// use robot go to capture screen
	imageBitmap := robotgo.CaptureScreen(initLoc.X, initLoc.Y, width, height)
	if imageBitmap == nil {
		log.Println("Invalid image. Skipping screen capture.")
		return
	}
	defer robotgo.FreeBitmap(imageBitmap)

	log.Println(imageBitmap)

	// save image?
	image := robotgo.ToImage(imageBitmap)

	// encode to png
	err := png.Encode(&s.imageBuffer, image)
	if err != nil {
		log.Fatalln("Image encoding failed:", err)
	}

	// write to file
	file, err := os.Create("img.jpeg")
	if err != nil {
		panic(err)
	}

	log.Println(file)

	fw := bufio.NewWriter(file)

	_, err = fw.Write(s.imageBuffer.Bytes())
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
