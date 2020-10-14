package point

type Point interface {
	// Raw gets the x, y, z fo the point
	Raw() *point

	Set2DCoordinate(x, y float32) Point
}

type point struct {
	X, Y, Z float32
}

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
