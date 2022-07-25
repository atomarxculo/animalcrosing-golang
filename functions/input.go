package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		con.PlayerMoving = true
		con.PlayerDir = 1
		con.PlayerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		con.PlayerMoving = true
		con.PlayerDir = 0
		con.PlayerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		con.PlayerMoving = true
		con.PlayerDir = 2
		con.PlayerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		con.PlayerMoving = true
		con.PlayerDir = 3
		con.PlayerRight = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		con.MusicPaused = !con.MusicPaused
	}
}
