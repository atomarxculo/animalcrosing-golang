package functions

import (
	con "animalcrosing/constans"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(con.ScreenWidth, con.ScreenHeigth, "Animal Crossing del chino")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	textures()

	con.PlayerSrc = rl.NewRectangle(0, 0, 48, 48)
	con.PlayerDest = rl.NewRectangle(200, 200, 60, 60)

	initMusic()

	con.Cam = rl.NewCamera2D(rl.NewVector2(float32(con.ScreenWidth/2), float32(con.ScreenHeigth/2)), rl.NewVector2(float32(con.PlayerDest.X-(con.PlayerDest.Width/2)), float32(con.PlayerDest.Y-(con.PlayerDest.Height/2))), 0.0, 1.5)

	con.Cam.Zoom = 3

	loadMap("two.map")
}

func loadMap(mapFile string) {
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	con.MapW = -1
	con.MapH = -1
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if con.MapW == -1 {
			con.MapW = m
		} else if con.MapH == -1 {
			con.MapH = m
		} else if i < con.MapW*con.MapH+2 {
			con.TileMap = append(con.TileMap, m)
		} else {
			con.SrcMap = append(con.SrcMap, sliced[i])
		}
	}

	if len(con.TileMap) > con.MapW*con.MapH {
		con.TileMap = con.TileMap[:len(con.TileMap)-1]
	}
}

func textures() {
	con.GrassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	con.HillSprite = rl.LoadTexture("res/Tilesets/Hills.png")
	con.FenceSprite = rl.LoadTexture("res/Tilesets/Fences.png")
	con.HouseSprite = rl.LoadTexture("res/Tilesets/Wooden House.png")
	con.WaterSprite = rl.LoadTexture("res/Tilesets/Water.png")
	con.TilledSprite = rl.LoadTexture("res/Tilesets/Tilled Dirt.png")
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
