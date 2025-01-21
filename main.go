package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	. "screen2text/app"

	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.design/x/mainthread"
)

func main() { mainthread.Init(fn) }

func fn() {
	InitClipboard()

	defer mainthread.Call(func() { glfw.Terminate() })

	hotkeyChan := RegisterHotkey()

	go func() {
		for {
			select {
			case <-hotkeyChan:
				fmt.Println("Hotkey pressed!")
				StartSelection()
			}
		}
	}()

	startTray, endTray := StartSystray()

	startTray()
	defer endTray()
	select {}
}
