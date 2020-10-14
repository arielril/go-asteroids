package object

import (
	"github.com/arielril/go-asteroids/point"
)

// Object is the representation of the master object
type Object interface {
	// Draw the object in the window
	Draw()

	// Get the raw values of the object
	Raw() *Struct

	RotateRight() Object
	RotateLeft() Object
}

// Data is the formatted data from the template file
type Data [][]int

// Struct is the representation of an Object
type Struct struct {
	Data Data

	Pos point.Point

	Rotation float32
}

// New creates a new object
func New(data Data, p point.Point) Object {
	if p == nil {
		p = point.New(0, 0, 0)
	}

	return &Struct{
		Data:     data,
		Pos:      p,
		Rotation: 0,
	}
}

// Raw gets the raw values from an Object
func (o *Struct) Raw() *Struct {
	return o
}

// Draw draws the Object in the OpenGL window
func (o *Struct) Draw() {}

func _rotate(angle float32, o *Struct) {
	o.Rotation += angle
}

// RotateRight rotate the object to the right
func (o *Struct) RotateRight() Object {
	return o
}

// RotateLeft rotate the object to the left
func (o *Struct) RotateLeft() Object {
	return o
}
