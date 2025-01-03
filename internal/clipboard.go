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

		//clipboard.Write(clipboard.FmtImage, s.imageBuffer.Bytes())
		extracted_text, err := Ocr(s)
		if err != nil {
			return
		}
		clipboard.Write(clipboard.FmtText, []byte(extracted_text))

	} else {
		log.Println("Image buffer empty. Skipping clipboard copy.")
	}
}
