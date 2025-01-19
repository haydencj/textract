package app

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
)

// image processing

func ReadImage(s *State) {
	activeLoc := s.SystemMouse.activeLoc
	initLoc := s.SystemMouse.initLoc

	width := abs(activeLoc.X - initLoc.X)
	height := abs(activeLoc.Y - initLoc.Y)

	log.Println("image size:", width, height)

	// use robot go to capture screen
	imageBitmap := robotgo.CaptureScreen(initLoc.X+1, initLoc.Y+1, width-1, height-1)
	if imageBitmap == nil {
		log.Println("Invalid image. Skipping screen capture.")
		return
	}
	defer robotgo.FreeBitmap(imageBitmap)

	// save image?
	image := robotgo.ToImage(imageBitmap)

	// encode to png
	err := png.Encode(&s.imageBuffer, image)
	if err != nil {
		log.Fatalln("Image encoding failed:", err)
	}

	// write to file
	file, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}

	fw := bufio.NewWriter(file)

	_, err = fw.Write(s.imageBuffer.Bytes())
	if err != nil {
		panic(err)
	}

	err = fw.Flush() // writes data from buffer to file
	if err != nil {
		panic(err)
	}
}

func Ocr(s *State) (string, error) {
	start := time.Now()

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImageFromBytes(s.imageBuffer.Bytes())
	text, err := client.Text()
	if err != nil {
		return "", fmt.Errorf("unable to convert file: %v", err)

	}

	end := time.Now()
	log.Println("extracted text:", text)
	log.Println("elapsed ocr time:", end.Sub(start))

	return text, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
