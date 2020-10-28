package game

import (
	"math/rand"

	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/object/bullet"
	"github.com/arielril/go-asteroids/object/ship"
	"github.com/arielril/go-asteroids/object/shooter"
	"github.com/arielril/go-asteroids/point"
	"github.com/arielril/go-asteroids/util"
)

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
		shipMapping = make(map[ship.Format]object.Data)
		shipMapping[ship.Ship1] = util.GetObjectDataFromFile("./templates/ship1.txt")
		shipMapping[ship.Ship2] = util.GetObjectDataFromFile("./templates/ship2.txt")
		shipMapping[ship.Ship3] = util.GetObjectDataFromFile("./templates/ship3.txt")
		ships = make([]ship.Ship, maxEnemyShips)
		_randomCreateEnemyShips()

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

func _randomCreateEnemyShips() {
	for i := range ships {
		s := rand.Intn(3)
		ships[i] = ship.New(
			shipMapping[ship.Format(s)],
			point.New(float32(rand.Int()%100)/10.0, float32(rand.Int()%105)/10.0, 0),
		)
	}
}
