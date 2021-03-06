package game

import (
	"fmt"
	"sync"

	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/object/bullet"
	"github.com/arielril/go-asteroids/object/ship"
	"github.com/arielril/go-asteroids/object/shooter"
	"github.com/arielril/go-asteroids/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const maxBullets int = 10
const maxEnemyShips = 7
const maxEnemyBullets = 10

var once sync.Once

var shooterObj shooter.Shooter
var shipMapping map[ship.Format]object.Data
var ships []ship.Ship
var lives []object.Object
var bullets []bullet.Bullet
var enemyBullets []bullet.Bullet

// Init the game
func Init() {
	fmt.Println("Game initialized...")
	fps = util.NewFps()
	_initializeObjects()
}

// Display the game
func Display(w *glfw.Window) {
	if len(lives)-1 <= 0 {
		fmt.Printf("\nEnd game, no lives!\n\n")
		w.SetShouldClose(true)
	}

	displayFps()
	displayScenario()
}
