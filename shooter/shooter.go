package shooter

import "github.com/arielril/go-asteroids/object"

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
	o := object.New(raw, nil)

	s := &shooter{
		o.Raw(),
	}

	return s
}

func (s *shooter) Move() {
	// TODO implement shooter move
}

func (s *shooter) Shoot() {
	// TODO implement shooter shoot
}
