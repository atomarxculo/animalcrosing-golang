package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Quit() {
	rl.UnloadTexture(con.GrassSprite)
	rl.UnloadTexture(con.PlayerSprite)
	rl.UnloadMusicStream(con.Music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
