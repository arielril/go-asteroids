package bullet

import (
	"fmt"

	"github.com/arielril/go-asteroids/object"
	"github.com/arielril/go-asteroids/point"
)

// Bullet is the interface of a bullet
type Bullet interface {
	object.Object
	Hit(o object.Object) bool
	ShouldRemove() bool
}

type bullet struct {
	*object.Struct
}

type xx struct {
	Enemy   string
	Shooter string
}

// Type of bullets
var Type *xx = &xx{
	Enemy:   "ENEMY",
	Shooter: "SHOOTER",
}

var objectData map[string]object.Data = make(map[string]object.Data)

// New creates a new bullet
func New(tp string, origin point.Point, rotation float32) Bullet {
	fmt.Printf("new bullet %v \n", objectData)

	o := object.New(
		objectData[tp],
		point.New(origin.Raw().X, origin.Raw().Y, origin.Raw().Z),
	)

	o.SetRotation(rotation)

	b := &bullet{
		Struct: o.Raw(),
	}

	return b
}

func (b *bullet) Hit(o object.Object) bool {
	return false
}

func (b *bullet) ShouldRemove() bool {
	position := b.Pos.Raw()

	// the bullet is between (X, Y) = ([0, 10], [0, 10])
	return position.X >= 10 ||
		position.X <= 0 ||
		position.Y >= 10 ||
		position.Y <= 0
}
