package constans

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 1280
	ScreenHeigth = 720
)

var (
	Running  = true
	BkgColor = rl.NewColor(147, 211, 196, 255)

	GrassSprite  rl.Texture2D
	HillSprite   rl.Texture2D
	FenceSprite  rl.Texture2D
	HouseSprite  rl.Texture2D
	WaterSprite  rl.Texture2D
	TilledSprite rl.Texture2D
	Tex          rl.Texture2D
	PlayerSprite rl.Texture2D

	PlayerSrc                                     rl.Rectangle
	PlayerDest                                    rl.Rectangle
	PlayerMoving                                  bool
	PlayerDir                                     int
	PlayerUp, PlayerDown, PlayerRight, PlayerLeft bool
	PlayerFrame                                   int

	FrameCount int

	TileDest   rl.Rectangle
	TileSrc    rl.Rectangle
	TileMap    []int
	SrcMap     []string
	MapW, MapH int

	PlayerSpeed float32 = 1.4

	MusicPaused bool
	Music       rl.Music

	Cam rl.Camera2D
)
