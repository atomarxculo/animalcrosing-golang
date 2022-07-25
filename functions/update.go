package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	con.Running = !rl.WindowShouldClose()

	con.PlayerSrc.X = con.PlayerSrc.Width * float32(con.PlayerFrame)
	if con.PlayerMoving {
		if con.PlayerUp {
			con.PlayerDest.Y -= con.PlayerSpeed
		}
		if con.PlayerDown {
			con.PlayerDest.Y += con.PlayerSpeed
		}
		if con.PlayerLeft {
			con.PlayerDest.X -= con.PlayerSpeed
		}
		if con.PlayerRight {
			con.PlayerDest.X += con.PlayerSpeed
		}
		if con.FrameCount%8 == 1 {
			con.PlayerFrame++
		}
	} else if con.FrameCount%45 == 1 {
		con.PlayerFrame++
	}

	con.FrameCount++
	if con.PlayerFrame > 3 {
		con.PlayerFrame = 0
	}
	if !con.PlayerMoving && con.PlayerFrame > 1 {
		con.PlayerFrame = 0
	}

	con.PlayerSrc.X = con.PlayerSrc.Width * float32(con.PlayerFrame)
	con.PlayerSrc.Y = con.PlayerSrc.Height * float32(con.PlayerDir)

	rl.UpdateMusicStream(con.Music)
	if con.MusicPaused {
		rl.PauseMusicStream(con.Music)
	} else {
		rl.ResumeMusicStream(con.Music)
	}

	con.Cam.Target = rl.NewVector2(float32(con.PlayerDest.X-(con.PlayerDest.Width/2)), float32(con.PlayerDest.Y-(con.PlayerDest.Height/2)))

	con.PlayerMoving = false
	con.PlayerUp, con.PlayerDown, con.PlayerRight, con.PlayerLeft = false, false, false, false
}
