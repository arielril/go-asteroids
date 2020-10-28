package shooter

import (
	"fmt"

	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/object/bullet"
	"github.com/arielril/go-asteroids/point"
)

// Shooter is the main character of the game
type Shooter interface {
	object.Object
	Move(m Movement)
	Shoot() bullet.Bullet
}

type shooter struct {
	*object.Struct
}

// Movement is the moves for the shooter
type Movement string

type moves struct {
	RotateLeft, RotateRight, Front Movement
}

// Moves is the set of moves available for the shooter
var Moves *moves = &moves{
	RotateLeft:  "ROTATE_LEFT",
	RotateRight: "ROTATE_RIGHT",
	Front:       "FRONT",
}

// NewFromRawObject create a shooter from the raw object.Data
func NewFromRawObject(raw object.Data) Shooter {
	o := object.New(raw, point.New(5, 5, 0))

	s := &shooter{
		Struct: o.Raw(),
	}

	return s
}

func (s *shooter) Move(m Movement) {
	// * when the shooter moves, rotate the OpenGL and update the direction vector
	switch m {
	case Moves.Front:
		s.DirectionVector.Set2DCoordinate(0, 1)
		break
	case Moves.RotateLeft:
		s.RotateLeft()
		break
	case Moves.RotateRight:
		s.RotateRight()
		break
	}
	fmt.Printf("Shooter pos (%v)\n", s.Pos.Raw())
}

func (s *shooter) Shoot() bullet.Bullet {
	b := bullet.New(
		bullet.Type.Shooter,
		s.Pos,
		s.Rotation,
	)

	fmt.Printf("shooting... %#v\n", b.Raw().Data)

	return b
}
