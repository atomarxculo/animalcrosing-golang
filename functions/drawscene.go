package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScene() {
	rl.DrawTexture(con.GrassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(con.PlayerSprite, con.PlayerSrc, con.PlayerDest, rl.NewVector2(con.PlayerDest.Width, con.PlayerDest.Height), 0, rl.White)
}
