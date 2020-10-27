package point

// Point is the interface of a Point
type Point interface {
	// Raw gets the x, y, z fo the point
	Raw() *point

	Set2DCoordinate(x, y float32) Point
	SetCoordinateFromPoint(p Point) Point
}

type point struct {
	X, Y, Z float32
}

// New creates a new Point
func New(x, y, z float32) Point {
	return &point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *point) Raw() *point {
	return p
}

func (p *point) Set2DCoordinate(x, y float32) Point {
	p.X = x
	p.Y = y

	return p
}

func (p *point) SetCoordinateFromPoint(p2 Point) Point {
	p2Raw := p2.Raw()
	p.X = p2Raw.X
	p.Y = p2Raw.Y
	p.Z = p2Raw.Z

	return p
}
