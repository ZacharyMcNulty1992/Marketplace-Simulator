package main

import (
	"actor"
	"fmt"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/window"
)

func main() {
	fmt.Printf("Initializing World\n")
	//TODO: create a world map
	createWindow()

	//TODO: change this so that it is creating a lot of new people.
	person := actor.NewPerson("merchant")
	//TODO: have the main thread be the thread that controls world events.
	person.Init()

}

func createWindow() {
	// Creates window and OpenGL context
	win, err := window.New("glfw", 800, 600, "Hello G3N", false)
	if err != nil {
		panic(err)
	}
	// Create OpenGL state
	gs, err := gls.New()
	if err != nil {
		panic(err)
	}

	// Sets the OpenGL viewport size the same as the window size
	// This normally should be updated if the window is resized.
	width, height := win.GetSize()
	gs.Viewport(0, 0, int32(width), int32(height))

	// Creates scene for 3D objects
	scene := core.NewNode()
}
