package boundingbox

import (
	"math"

	"github.com/arielril/go-asteroids/point"
)

// Struct is the bounding box data struct
type Struct struct {
	CenterPoint           point.Point
	HalfWidth, HalfHeight float64
}

// BoundingBox interface
type BoundingBox interface {
	Raw() *Struct
	HasCollidedWith(b BoundingBox) bool
}

// New create a new bounding box representation
func New(height, width float64, p point.Point) BoundingBox {
	halfHeight := (height * .07) / 2
	halfWidth := (width * .07) / 2

	centerPoint := point.New(
		p.Raw().X,
		p.Raw().Y,
		p.Raw().Z,
	)

	return &Struct{
		CenterPoint: centerPoint,
		HalfHeight:  halfHeight,
		HalfWidth:   halfWidth,
	}
}

// Raw returns the BoundingBox.Struct
func (b *Struct) Raw() *Struct {
	return b
}

// HasCollidedWith returns if two bounding boxes has collided
func (b *Struct) HasCollidedWith(bb BoundingBox) bool {
	bCenterRaw := b.CenterPoint.Raw()
	bbCenterRaw := bb.Raw().CenterPoint.Raw()

	xAbs := math.Abs(float64(bCenterRaw.X - bbCenterRaw.X))
	widthSum := b.HalfWidth + bb.Raw().HalfWidth

	if xAbs > widthSum {
		return false
	}

	yAbs := math.Abs(float64(bCenterRaw.Y - bbCenterRaw.Y))
	heightSum := b.HalfHeight + bb.Raw().HalfHeight

	if yAbs > heightSum {
		return false
	}

	return true
}
