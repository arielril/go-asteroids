package ship

import (
	"fmt"
	"math"
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
	RawShip() *Struct
}

// Struct is the representation of a ship
type Struct struct {
	*object.Struct

	TrajectoryProgress float32
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

	s := &Struct{
		Struct:             o.Raw(),
		TrajectoryProgress: 0,
	}

	return s
}

// Shoot from the ship
func (s *Struct) Shoot() (bullet.Bullet, bool) {
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

func _getTrajectoryNextPoint() point.Point {
	rand.Seed(time.Now().UnixNano())

	return point.New(
		rand.Float32()*100,
		rand.Float32()*100,
		0,
	)
}

func _vecMultiply(num float32, vec point.Point) point.Point {
	return point.New(
		vec.Raw().X*num,
		vec.Raw().Y*num,
		vec.Raw().Z*num,
	)
}

func _vecSum(v1, vec point.Point) point.Point {
	return point.New(
		vec.Raw().X+v1.Raw().X,
		vec.Raw().Y+v1.Raw().Y,
		vec.Raw().Z+v1.Raw().Z,
	)
}

// CreateTrajectory of the Bezier
func (s *Struct) CreateTrajectory() {
	//  C1(t) = (1-t)^2 * P0 + 2 * (1-t) * t * P1 + t^2 * P2
	rand.Seed(time.Now().UnixNano())

	p0 := s.Pos
	p1 := _vecSum(p0, point.New(rand.Float32(), rand.Float32(), 0))
	p2 := _getTrajectoryNextPoint()

	var t float32 = s.TrajectoryProgress

	m1 := _vecMultiply(
		float32(math.Pow(float64(1-t), 2)),
		p0,
	)
	m2 := _vecMultiply(
		2*(1-t)*t,
		p1,
	)
	m3 := _vecMultiply(
		float32(math.Pow(float64(t), 2)),
		p2,
	)

	c1 := _vecSum(
		_vecSum(m1, m2),
		m3,
	)

	fmt.Printf("new dir vec %v\n", c1)
	s.Pos = c1
}

// Hit is
func (s *Struct) Hit() {}

// RawShip returns the ship raw struct
func (s *Struct) RawShip() *Struct {
	return s
}
