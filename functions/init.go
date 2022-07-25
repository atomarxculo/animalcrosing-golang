package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(con.ScreenWidth, con.ScreenHeigth, "Animal Crossing del chino")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	textures()

	con.PlayerSrc = rl.NewRectangle(0, 0, 48, 48)
	con.PlayerDest = rl.NewRectangle(200, 200, 100, 100)

	initMusic()

	con.Cam = rl.NewCamera2D(rl.NewVector2(float32(con.ScreenWidth/2), float32(con.ScreenHeigth/2)), rl.NewVector2(float32(con.PlayerDest.X-(con.PlayerDest.Width/2)), float32(con.PlayerDest.Y-(con.PlayerDest.Height/2))), 0.0, 1.5)

	loadMap()
}

func loadMap() {
	con.MapW = 5
	con.MapH = 5
	for i := 0; i < (con.MapW * con.MapH); i++ {
		con.TileMap = append(con.TileMap, 1)
	}
}

func textures() {
	con.GrassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	con.TileDest = rl.NewRectangle(0, 0, 16, 16)
	con.TileSrc = rl.NewRectangle(0, 0, 16, 16)
	con.PlayerSprite = rl.LoadTexture("res/Characters/character.png")
}

func initMusic() {
	rl.InitAudioDevice()
	con.Music = rl.LoadMusicStream("res/musicbkgfarm.mp3")
	con.MusicPaused = false
	rl.PlayMusicStream(con.Music)
}
