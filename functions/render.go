package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.BeginDrawing()

	rl.ClearBackground(con.BkgColor)
	rl.BeginMode2D(con.Cam)

	DrawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}
