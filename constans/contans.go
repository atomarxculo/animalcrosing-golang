package constans

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 1000
	ScreenHeigth = 480
)

var (
	Running  = true
	BkgColor = rl.NewColor(147, 211, 196, 255)

	GrassSprite  rl.Texture2D
	PlayerSprite rl.Texture2D

	PlayerSrc                                     rl.Rectangle
	PlayerDest                                    rl.Rectangle
	PlayerMoving                                  bool
	PlayerDir                                     int
	PlayerUp, PlayerDown, PlayerRight, PlayerLeft bool
	PlayerFrame                                   int

	FrameCount int

	PlayerSpeed float32 = 3

	MusicPaused bool
	Music       rl.Music

	Cam rl.Camera2D
)
