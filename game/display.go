package game

import (
	"fmt"

	"github.com/arielril/go-asteroids/object/bullet"
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
	bulletList := make([]bullet.Bullet, 0)
	for _, b := range bullets {
		b.Draw()
		b.Raw().DirectionVector.Set2DCoordinate(0, 1)
		if !b.ShouldRemove() {
			bulletList = append(bulletList, b)
		}
	}
	bullets = bulletList

	// Ships
	for _, s := range ships {
		s.Shoot()
		s.Draw()
	}
	enemyBulletsList := make([]bullet.Bullet, 0)
	for _, b := range enemyBullets {
		if b == nil {
			continue
		}

		b.Draw()
		b.Raw().DirectionVector.Set2DCoordinate(0, 1)
		if !b.ShouldRemove() {
			enemyBulletsList = append(enemyBulletsList, b)
		}
	}
	enemyBullets = enemyBulletsList

	if shooterObj != nil {
		// Shooter
		shooterObj.Draw()
	}
}
