package game

import (
	"fmt"
	"sync"

	"github.com/arielril/go-asteroids/ship"
	"github.com/arielril/go-asteroids/shooter"
	"github.com/arielril/go-asteroids/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var once sync.Once

var shooterObj shooter.Shooter
var ships map[ship.Format]ship.Ship
var lives []interface{}

// Init the game
func Init() {
	fmt.Println("Game initialized...")
	fps = util.NewFps()
}

func _initializeObjects() {
	once.Do(func() {
		// Shooter
		shooterObj = shooter.NewFromRawObject(
			util.GetObjectDataFromFile("../templates/shooter.txt"),
		)

		// Ships
		ships[ship.Ship1] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("../templates/ship1.txt"),
		)
		ships[ship.Ship2] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("../templates/ship2.txt"),
		)
		ships[ship.Ship3] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("../templates/ship3.txt"),
		)

		// Bullets
	})
}

// Display the game
func Display(w *glfw.Window) {
	displayFps()
}
