package game

import (
	"fmt"

	"github.com/arielril/go-asteroids/object/bullet"
	"github.com/arielril/go-asteroids/object/ship"
	"github.com/arielril/go-asteroids/point"
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
		b.UpdateBoundingBox()
		if !b.ShouldRemove() {
			bulletList = append(bulletList, b)
		}
	}
	bullets = bulletList

	// Ships
	for _, s := range ships {
		enemyShoot(s)
		s.UpdateBoundingBox()
		s.Draw()
		s.Raw().DirectionVector = point.New(0, .3, 0)
		// if s.RawShip().TrajectoryProgress <= 0 ||
		// 	s.RawShip().TrajectoryProgress >= 1 {
		// 	s.RawShip().TrajectoryProgress = 0
		// }
		// s.CreateTrajectory()
		// s.RawShip().TrajectoryProgress += 0.0001
	}
	enemyBulletsList := make([]bullet.Bullet, 0)
	for _, b := range enemyBullets {
		if b == nil {
			continue
		}

		b.Draw()
		b.Raw().DirectionVector.Set2DCoordinate(0, 1)
		b.UpdateBoundingBox()
		if !b.ShouldRemove() {
			enemyBulletsList = append(enemyBulletsList, b)
		}
	}
	enemyBullets = enemyBulletsList

	if shooterObj != nil {
		// Shooter
		shooterObj.Draw()
		shooterObj.UpdateBoundingBox()
	}

	for _, b := range bullets {
		for _, s := range ships {
			if b.Raw().BoundingBox.HasCollidedWith(s.Raw().BoundingBox) {
				_killShip(s)
				_removeBullet(b)
			}
		}
	}

	for _, s := range ships {
		if s.Raw().BoundingBox.HasCollidedWith(shooterObj.Raw().BoundingBox) {
			_removeLifeFromShooter()
		}
	}

	for _, b := range enemyBulletsList {
		if b.Raw().BoundingBox.HasCollidedWith(shooterObj.Raw().BoundingBox) {
			_removeLifeFromShooter()
			_removeEnemyBullet(b)
		}
	}
}

func _killShip(shipToKill ship.Ship) {
	shipList := make([]ship.Ship, 0)

	for _, s := range ships {
		if shipToKill == s {
			continue
		}
		shipList = append(shipList, s)
	}

	ships = shipList
}

func _removeBullet(bulletToRemove bullet.Bullet) {
	bulletList := make([]bullet.Bullet, 0)

	for _, b := range bullets {
		if bulletToRemove == b {
			continue
		}
		bulletList = append(bulletList, b)
	}

	bullets = bulletList
}

func _removeEnemyBullet(bulletToRemove bullet.Bullet) {
	bulletList := make([]bullet.Bullet, 0)

	for _, b := range enemyBullets {
		if bulletToRemove == b {
			continue
		}
		bulletList = append(bulletList, b)
	}

	enemyBullets = bulletList
}

func _removeLifeFromShooter() {
//	lives = lives[:len(lives)-1]
}
