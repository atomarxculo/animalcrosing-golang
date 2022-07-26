package functions

import (
	con "animalcrosing/constans"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScene() {
	//rl.DrawTexture(con.GrassSprite, 100, 50, rl.White)

	for i := 0; i < len(con.TileMap); i++ {
		if con.TileMap[i] != 0 {
			con.TileDest.X = con.TileDest.Width * float32(i%con.MapW)
			con.TileDest.Y = con.TileDest.Height * float32(i/con.MapW)

			if con.SrcMap[i] == "g" {
				con.Tex = con.GrassSprite
			}
			if con.SrcMap[i] == "l" {
				con.Tex = con.HillSprite
			}
			if con.SrcMap[i] == "f" {
				con.Tex = con.FenceSprite
			}
			if con.SrcMap[i] == "h" {
				con.Tex = con.HouseSprite
			}
			if con.SrcMap[i] == "w" {
				con.Tex = con.WaterSprite
			}
			if con.SrcMap[i] == "t" {
				con.Tex = con.TilledSprite
			}

			if con.SrcMap[i] == "h" || con.SrcMap[i] == "f" {
				con.TileSrc.X = 0
				con.TileSrc.Y = 0
				rl.DrawTexturePro(con.GrassSprite, con.TileSrc, con.TileDest, rl.NewVector2(con.TileDest.Width, con.TileDest.Height), 0, rl.White)
			}

			con.TileSrc.X = con.TileSrc.Width * float32((con.TileMap[i]-1)%int(con.Tex.Width/int32(con.TileSrc.Width)))
			con.TileSrc.Y = con.TileSrc.Height * float32((con.TileMap[i]-1)/int(con.Tex.Width/int32(con.TileSrc.Width)))
			rl.DrawTexturePro(con.Tex, con.TileSrc, con.TileDest, rl.NewVector2(con.TileDest.Width, con.TileDest.Height), 0, rl.White)
		}
	}

	rl.DrawTexturePro(con.PlayerSprite, con.PlayerSrc, con.PlayerDest, rl.NewVector2(con.PlayerDest.Width, con.PlayerDest.Height), 0, rl.White)
}
