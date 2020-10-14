package object

import (
	"github.com/arielril/go-asteroids/point"
)

type Object interface {
	// Draw the object in the window
	Draw()

	// Get the raw values of the object
	Raw() *object

	RotateRight() Object
	RotateLeft() Object
}

type object struct {
	Data [][]int

	Pos point.Point

	Rotation float32
}

func New() Object {
	return &object{
		Pos:      point.New(0, 0, 0),
		Rotation: 0,
	}
}

func (o *object) Raw() *object {
	return o
}

func (o *object) Draw() {}

func _rotate(angle float32, o *object) {
	o.Rotation += angle
}

func (o *object) RotateRight() Object {
	return o
}

func (o *object) RotateLeft() Object {
	return o
}
