package game

import (
	"fmt"

	"github.com/arielril/go-asteroids/util"
)

var fps util.FPS

func displayFps() {
	acc := fps.SetFPS().GetAccumulated()
	if acc >= 1 { // print every second
		fmt.Printf("FPS: %v\n", fps.GetFPS())
		fps.Reset()
	}
}
