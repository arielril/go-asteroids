package bullet

import "github.com/arielril/go-asteroids/object"

// Bullet is the interface of a bullet
type Bullet interface {
	object.Object
	Hit(o object.Object) bool
}

type bullet struct {
	*object.Struct
}

func (b *bullet) Hit(o object.Object) bool {
	return false
}
