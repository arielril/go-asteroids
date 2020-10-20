package object

import (
	"github.com/arielril/go-asteroids/point"
	"github.com/go-gl/gl/v2.1/gl"
)

var colors Data

// Object is the representation of the master object
type Object interface {
	// Draw the object in the window
	Draw()

	// Get the raw values of the object
	Raw() *Struct

	RotateRight() Object
	RotateLeft() Object

	CreateBoundingBox(o Object) interface{}
}

// Data is the formatted data from the template file
type Data [][]int

// Struct is the representation of an Object
type Struct struct {
	Data Data

	Pos point.Point

	Rotation float32

	BoundingBox interface{}
}

// New creates a new object
func New(data Data, p point.Point) Object {
	if p == nil {
		p = point.New(0, 0, 0)
	}

	return &Struct{
		Data:        data,
		Pos:         p,
		Rotation:    0,
		BoundingBox: nil,
	}
}

// Raw gets the raw values from an Object
func (o *Struct) Raw() *Struct {
	return o
}

// SetColors initialize the colors of the game
func SetColors(gameColors Data) {
	colors = gameColors
}

func _setColor(templateColor int) {
	for _, c := range colors {
		if c[0] == templateColor {
			r := float32(c[1])
			g := float32(c[2])
			b := float32(c[3])

			gl.Color3f(r, g, b)
		}
	}
}

// Draw draws the Object in the OpenGL window
func (o *Struct) Draw() {
	objHeight := float32(len(o.Data))

	gl.PushMatrix()
	{
		gl.Translatef(o.Pos.Raw().X, o.Pos.Raw().Y, 0)
		gl.Scalef(.07, .07, 0)

		for i := range o.Data {
			y := float32(i)

			for j, color := range o.Data[i] {
				_setColor(color)

				x := float32(j)

				gl.Begin(gl.QUADS)
				{
					gl.Vertex2f(x, objHeight-y)
					gl.Vertex2f(x+1, objHeight-y)
					gl.Vertex2f(x+1, objHeight-y-1)
					gl.Vertex2f(x, objHeight-y-1)
				}
				gl.End()
			}
		}
	}
	gl.PopMatrix()
}

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

// CreateBoundingBox create a bounding box for each object and to test collisions
func (o *Struct) CreateBoundingBox(obj Object) interface{} {
	return nil
}
