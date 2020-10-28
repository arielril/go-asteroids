package ship

import (
	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/point"
)

// Ship is the representation of a ship
type Ship interface {
	object.Object
	Shoot()
	CreateTrajectory()
	Hit()
}

type ship struct {
	*object.Struct
}

// Format of the ships
type Format int

const (
	// Ship1 is the first format
	Ship1 Format = iota

	// Ship2 is the second format
	Ship2

	// Ship3 is the third format
	Ship3
)

// New creates a new Ship from raw object.Data
func New(raw object.Data, p point.Point) Ship {
	o := object.New(raw, p)

	s := &ship{
		o.Raw(),
	}

	return s
}

func (s *ship) Shoot() {
	// TODO implement ship shoot
}

func (s *ship) CreateTrajectory() {
	// TODO implement ship trajectory
}

func (s *ship) Hit() {}
