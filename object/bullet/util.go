package bullet

import "github.com/arielril/go-asteroids/object"

// SetObjectData initialize the object data
func SetObjectData(od object.Data, tp string) {
	objectData[tp] = od
}
