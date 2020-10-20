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

func displayScenario() {
	// Lives
	for _, l := range lives {
		l.Draw()
	}

	// Bullets
	for _, b := range bullets {
		b.Draw()
	}

	// Ships
	for _, s := range ships {
		s.Draw()
	}

	if shooterObj != nil {
		// Shooter
		shooterObj.Draw()
	}
}
