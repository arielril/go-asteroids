package ship

import (
	"math/rand"
	"time"

	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/object/bullet"
	"github.com/arielril/go-asteroids/point"
)

// Ship is the representation of a ship
type Ship interface {
	object.Object
	Shoot() (bullet.Bullet, bool)
	CreateTrajectory()
	Hit()
}

type ship struct {
	*object.Struct

	trajectoryProgress float64
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
	rand.Seed(time.Now().UnixNano())

	o := object.New(raw, p)

	s := &ship{
		Struct:             o.Raw(),
		trajectoryProgress: 0,
	}

	return s
}

func (s *ship) Shoot() (bullet.Bullet, bool) {
	r := rand.Uint32() % 5000

	if r >= 0 && r <= 100 {
		b := bullet.New(
			bullet.Type.Enemy,
			s.Pos,
			s.Rotation,
		)
		return b, true
	}

	return nil, false
}

func (s *ship) CreateTrajectory() {
	// TODO implement ship trajectory
	//  C1(t) = (1-t)2 * P0 + 2 * (1-t) * t * P1 + t2 * P2
}

func (s *ship) Hit() {}
