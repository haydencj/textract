package main

// TODO: #1 Remove wildcard import for internal package.
import (
	"fmt"
	. "screen2text/app"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// func main() {
// 	mainthread.Init(fn)
// }

func main() {
	InitClipboard()

	fmt.Println("test 1")

	var err error
	// mainthread.Call(func() {
	// 	err = glfw.Init()
	// })

	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	win, err := NewWindow()
	if err != nil {
		panic(err)
	}
	startTray, endTray := StartSystray()
	startTray()
	// mainthread.Call(func() { startTray() })

	win.Run()
	fmt.Println("test 2")

	defer Terminate()

	//SwitchUI()

	fmt.Println("test 3")
	endTray()
	// window, err := NewWindow()
	// if err != nil {
	// 	panic(err)
	// }
	// // runs every frame
	// window.Run()

}
