package internal

import (
	"log"

	"golang.design/x/clipboard"
)

func InitClipboard() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		log.Fatalln("Clipboard package not ready for use:", err)
	}
}

func Copy(s *State) {
	if s.imageBuffer.Len() != 0 {

		clipboard.Write(clipboard.FmtImage, s.imageBuffer.Bytes())

	} else {
		log.Println("Image buffer empty. Skipping clipboard copy.")
	}
}
