package game

import (
	"fmt"
	"sync"

	"github.com/arielril/go-asteroids/bullet"
	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/point"
	"github.com/arielril/go-asteroids/ship"
	"github.com/arielril/go-asteroids/shooter"
	"github.com/arielril/go-asteroids/util"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const maxBullets int = 10

var once sync.Once

var shooterObj shooter.Shooter
var shipMapping map[ship.Format]ship.Ship
var ships []ship.Ship
var lives []object.Object
var bullets []bullet.Bullet

// Init the game
func Init() {
	fmt.Println("Game initialized...")
	fps = util.NewFps()
	_initializeObjects()
}

func _initializeObjects() {
	once.Do(func() {
		// Colors
		colors := util.GetObjectDataFromFile("./templates/color.txt")
		object.SetColors(colors)

		// Shooter
		shooterObj = shooter.NewFromRawObject(
			util.GetObjectDataFromFile("./templates/shooter.txt"),
		)

		// Ships
		shipMapping = make(map[ship.Format]ship.Ship)
		shipMapping[ship.Ship1] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("./templates/ship1.txt"),
		)
		shipMapping[ship.Ship2] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("./templates/ship2.txt"),
		)
		shipMapping[ship.Ship3] = ship.NewFromRawObject(
			util.GetObjectDataFromFile("./templates/ship3.txt"),
		)

		// Lives
		lifeTemplate := util.GetObjectDataFromFile("./templates/life.txt")
		lives = append(
			lives,
			// the player start with 3 lives
			object.New(lifeTemplate, point.New(9.5, 9.5, 0)),
			object.New(lifeTemplate, point.New(9, 9.5, 0)),
			object.New(lifeTemplate, point.New(8.5, 9.5, 0)),
		)

		// Bullet data objects
		bullet.SetObjectData(
			util.GetObjectDataFromFile("./templates/bullet_enemy.txt"),
			bullet.Type.Enemy,
		)
		bullet.SetObjectData(
			util.GetObjectDataFromFile("./templates/bullet_shooter.txt"),
			bullet.Type.Shooter,
		)
	})
}

// Display the game
func Display(w *glfw.Window) {
	displayFps()
	displayScenario()
}

// MoveShooter moves the shooter object
func MoveShooter(m shooter.Movement) {
	shooterObj.Move(m)
}

// ShooterShoot make the shooter shoot
func ShooterShoot() {
	if len(bullets)+1 >= maxBullets {
		return
	}

	shoot := shooterObj.Shoot()
	bullets = append(bullets, shoot)
}
