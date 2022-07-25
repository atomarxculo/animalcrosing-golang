package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(con.ScreenWidth, con.ScreenHeigth, "Animal Crossing del chino")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	con.GrassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	con.PlayerSprite = rl.LoadTexture("res/Characters/character.png")

	con.PlayerSrc = rl.NewRectangle(0, 0, 48, 48)
	con.PlayerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	con.Music = rl.LoadMusicStream("res/musicbkgfarm.mp3")
	con.MusicPaused = false
	rl.PlayMusicStream(con.Music)

	con.Cam = rl.NewCamera2D(rl.NewVector2(float32(con.ScreenWidth/2), float32(con.ScreenHeigth/2)), rl.NewVector2(float32(con.PlayerDest.X-(con.PlayerDest.Width/2)), float32(con.PlayerDest.Y-(con.PlayerDest.Height/2))), 0.0, 2.0)
}
