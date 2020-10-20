package shooter

import (
	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/point"
)

// Shooter is the main character of the game
type Shooter interface {
	object.Object
	Move()
	Shoot()
}

type shooter struct {
	*object.Struct
}

// NewFromRawObject create a shooter from the raw object.Data
func NewFromRawObject(raw object.Data) Shooter {
	o := object.New(raw, point.New(5, 5, 0))

	s := &shooter{
		o.Raw(),
	}

	return s
}

func _updateDirection(s *shooter) {}

func (s *shooter) Move() {
	// TODO implement shooter move

	// * when the shooter moves, rotate the OpenGL and update the direction vector

	/*
		1. glRotate(ang)
		2. glTranslate(0,1,0)
		3. draw shooter
		4. update direction
	*/
}

func (s *shooter) Shoot() {
	// TODO implement shooter shoot
}
