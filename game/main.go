package game

import (
	"fmt"

	"github.com/arielril/go-asteroids/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Init the game
func Init() {
	fmt.Println("Game initialized...")
	fps = util.NewFps()
}

// Display the game
func Display(w *glfw.Window) {
	displayFps()
}
