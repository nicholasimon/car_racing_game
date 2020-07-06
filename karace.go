package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ( // MARK: var
	// imgs
	car1      = rl.NewRectangle(193, 18, 32, 25)
	car1l     = rl.NewRectangle(225, 18, 33, 27)
	car1d     = rl.NewRectangle(260, 14, 34, 33)
	car1ul    = rl.NewRectangle(367, 15, 34, 31)
	car1dl    = rl.NewRectangle(332, 14, 34, 34)
	car1u     = rl.NewRectangle(297, 16, 34, 31)
	smokeur   = rl.NewRectangle(209, 57, 18, 12)
	smoke     = rl.NewRectangle(194, 46, 17, 9)
	smoker    = rl.NewRectangle(249, 46, 18, 9)
	tracktop1 = rl.NewRectangle(208, 0, 16, 16)
	track1    = rl.NewRectangle(192, 0, 16, 16)
	grass1    = rl.NewRectangle(0, 0, 16, 16)
	tree1     = rl.NewRectangle(176, 111, 16, 16)
	imgs      rl.Texture2D
	// extras
	extrasmap = make([]string, drawmapa)
	trackmap  = make([]string, drawmapa)
	// player
	player, playerh, playerv int
	playerdir                = 2
	// other
	pause bool
	// draw map
	scale                                                    = 5
	scalea                                                   = scale * scale
	drawblock, drawblocknext, drawblocknexth, drawblocknextv int
	drawmapw                                                 = layoutw * scale
	drawmaph                                                 = layouth * scale
	drawmapa                                                 = drawmaph * drawmapw
	drawmap                                                  = make([]string, drawmapa)
	// layout map
	startblockv, startblockh                int
	layoutw                                 = 64
	layouth                                 = 36
	layouta                                 = layoutw * layouth
	layoutmap                               = make([]string, layouta)
	maindir, dir, dirend                    int
	layoutblock, layoutblockh, layoutblockv int
	trackl                                  int
	trackmaxl                               = 4
	trackminl                               = 1
	minb                                    = trackminl + 1
	maxb                                    = trackmaxl + 1
	// core
	lockview                             bool
	screenw, screenh, screena            int
	monh32, monw32                       int32
	monitorh, monitorw, monitornum       int
	grid16on, grid4on, debugon, lrg, sml bool
	framecount                           int
	mousepos                             rl.Vector2
	camera                               rl.Camera2D
)

func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "karace")
	setscreen()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "karace")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(30)
	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		// MARK: draw map layer 1 / left up

		count := 0
		drawx := int32(0)
		drawy := int32(0)
		drawblock = drawblocknext
		for a := 0; a < screena; a++ {

			checkdrawmap := drawmap[drawblock]
			checkextras := extrasmap[drawblock]
			checktrack := trackmap[drawblock]

			switch checkdrawmap {
			case ".":
				//	rl.DrawRectangle(drawx, drawy, 15, 15, rl.Fade(rl.Green, 0.4))
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, grass1, v2, rl.White)
			case " ":
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Black)
			}
			switch checkextras {
			case "tree1", "tree2", "tree3", "tree4", "tree5", "tree6":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, tree1, v2, rl.White)
				//	rl.DrawRectangle(drawx, drawy, 15, 15, rl.Fade(rl.DarkGreen, 0.8))
			}
			switch checktrack {
			case "t":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, track1, v2, rl.White)

			case "tt":
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, tracktop1, v2, rl.White)
			}

			drawblock++
			count++
			drawx += 16
			if count == screenw {
				count = 0
				drawx = 0
				drawy += 16
				drawblock += drawmapw - screenw
			}

		}

		// MARK: draw map layer2
		count = 0
		drawx = int32(0)
		drawy = int32(0)
		drawblock = drawblocknext
		for a := 0; a < screena; a++ {
			// draw player
			if drawblock == player {
				//	rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
				v2 := rl.NewVector2(float32(drawx), float32(drawy))
				switch playerdir {
				case 1:
					v2smoke := rl.NewVector2(float32(drawx-5), float32(drawy+20))
					rl.DrawTextureRec(imgs, smokeur, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1u, v2, rl.White)
				case 2:
					v2smoke := rl.NewVector2(float32(drawx-12), float32(drawy+8))
					rl.DrawTextureRec(imgs, smoke, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1, v2, rl.White)
				case 3:
					v2smoke := rl.NewVector2(float32(drawx-12), float32(drawy+8))
					rl.DrawTextureRec(imgs, smoke, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1d, v2, rl.White)
				case 4:
					v2smoke := rl.NewVector2(float32(drawx+26), float32(drawy+8))
					rl.DrawTextureRec(imgs, smoker, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1l, v2, rl.White)
				case 5:
					v2smoke := rl.NewVector2(float32(drawx+26), float32(drawy+8))
					rl.DrawTextureRec(imgs, smoker, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1dl, v2, rl.White)
				case 6:
					v2smoke := rl.NewVector2(float32(drawx+26), float32(drawy+8))
					rl.DrawTextureRec(imgs, smoker, v2smoke, rl.White)
					rl.DrawTextureRec(imgs, car1ul, v2, rl.White)
				}
			}
			drawblock++
			count++
			drawx += 16
			if count == screenw {
				count = 0
				drawx = 0
				drawy += 16
				drawblock += drawmapw - screenw
			}

		}

		rl.EndMode2D() // MARK: draw no camera
		fx()
		if debugon {
			debug()
		}

		rl.EndDrawing()

		input()
		updateall()

	}
	rl.CloseWindow()
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	createlevel()
	pconsole()

	raylib()
}
func updateall() { // MARK: updateall()
	upcamera()
	animations()
}
func fx() { // MARK: fx()

	scany := int32(0)
	for a := 0; a < monitorh; a++ {
		rl.DrawLine(0, scany, monw32, scany, rl.Fade(rl.Black, 0.4))
		scany += 2
		a++

	}

}
func animations() { // MARK: animations()

	if framecount%3 == 0 {
		smoke.X += 17
		smoker.X -= 18
		smokeur.X -= 13
		smokeur.Y += 11
		if car1.Y == 18 {
			car1.Y = 19
		} else if car1.Y == 19 {
			car1.Y = 18
		}
		if car1l.Y == 18 {
			car1l.Y = 19
		} else if car1l.Y == 19 {
			car1l.Y = 18
		}
		if car1d.Y == 14 {
			car1d.Y = 15
		} else if car1d.Y == 15 {
			car1d.Y = 14
		}
		if car1ul.Y == 15 {
			car1ul.Y = 16
		} else if car1ul.Y == 16 {
			car1ul.Y = 15
		}
		if car1dl.Y == 14 {
			car1dl.Y = 15
		} else if car1dl.Y == 15 {
			car1dl.Y = 14
		}
		if car1u.Y == 16 {
			car1u.Y = 17
		} else if car1u.Y == 17 {
			car1u.Y = 16
		}
	}
	if smoke.X > 212 {
		smoke.X = 194
	}
	if smoker.X < 230 {
		smoker.X = 249
	}
	if smokeur.X < 195 {
		smokeur.X = 209
		smokeur.Y = 57
	}

}
func upcamera() { // MARK: camera()

	horizvert()
	if camera.Zoom == 1.0 {
		if playerv > drawblocknextv+(screenw/2) {
			if drawblocknextv < drawmapw-screenw {
				drawblocknext++
			}
		} else if playerv < drawblocknextv+(screenw/2) && playerv > screenw/2 {
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}
		if playerh > drawblocknexth+(screenh/2) {
			if drawblocknexth < (drawmaph - screenh) {
				drawblocknext += drawmapw
			}
		} else if playerh < drawblocknexth+(screenh/2) {
			if drawblocknexth > 0 {
				drawblocknext -= drawmapw
			}
		}
	}
	if camera.Zoom == 2.0 {
		if playerh > drawblocknexth+(screenh/4) {
			if drawblocknexth < (drawmaph - screenh) {
				drawblocknext += drawmapw
			}
		} else if playerh < drawblocknexth+(screenh/4) {
			if drawblocknexth > 0 {
				drawblocknext -= drawmapw
			}
		}
		if playerv > drawblocknextv+(screenw/4) {
			if drawblocknextv < drawmapw-screenw {
				drawblocknext++
			}
		} else if playerv < drawblocknextv+(screenw/4) && playerv > screenw/4 {
			if drawblocknextv > 0 {
				drawblocknext--
			}
		}

		if playerv > drawmapw-(((screenw/4)*3)+1) {
			if rl.IsKeyDown(rl.KeyRight) {
				if camera.Target.X < float32((monitorw/2)-16) {
					camera.Target.X += 16
				}
			}
		}
		if playerv < drawmapw-(screenw/4+2) {
			if camera.Target.X != 0 {
				if rl.IsKeyDown(rl.KeyLeft) {
					camera.Target.X -= 16
				}
				if camera.Target.X <= 0 {
					camera.Target.X = 0
				}
			}
		}
		if playerh > drawmaph-((screenh/4)*3) {
			if rl.IsKeyDown(rl.KeyDown) {
				if camera.Target.Y < float32((monitorh/2)-16) {
					camera.Target.Y += 16
				}
			}
		}
		if playerh < drawmaph-(screenh/4+1) {
			if camera.Target.Y != 0 {
				if rl.IsKeyDown(rl.KeyUp) {
					camera.Target.Y -= 16
				}
				if camera.Target.Y <= 0 {
					camera.Target.Y = 0
				}
			}
		}

	}
}
func tracktiles() { // MARK: tracktiles()

	for a := 0; a < drawmapa; a++ {
		if drawmap[a] == " " {
			if drawmap[a-drawmapw] == " " {
				trackmap[a] = "t"
			} else if drawmap[a-drawmapw] == "." {
				trackmap[a] = "tt"
			}
		}
	}

}
func createextras() { // MARK: camera()
	for a := 0; a < drawmapa; a++ {
		if rolldice()+rolldice()+rolldice() >= 17 {
			if drawmap[a] == "." {
				switch rolldice() {
				case 1:
					extrasmap[a] = "tree1"
				case 2:
					extrasmap[a] = "tree2"
				case 3:
					extrasmap[a] = "tree3"
				case 4:
					extrasmap[a] = "tree4"
				case 5:
					extrasmap[a] = "tree5"
				case 6:
					extrasmap[a] = "tree6"
				}
			}
		}
	}
}
func createlevel() { // MARK: createlevel()
	for a := 0; a < layouta; a++ {
		layoutmap[a] = "."
	}

	layoutblock = trackminl + 1
	layoutblock += (trackminl + 1) * layoutw
	layoutmap[layoutblock] = " "

	player = (trackminl + 2) * scale
	player += (((trackminl + 1) * scale) * drawmapw) + (drawmapw * 2)
	horizvert()

	for {
		horizvert()

		if layoutblockv < layoutw-trackmaxl {

			next := rInt(trackminl, trackmaxl)
			if layoutblockh < trackmaxl {
				if flipcoin() {
					for a := 0; a < next; a++ {
						horizvert()
						if layoutblockv > layoutw-trackmaxl {
							break
						}
						layoutmap[layoutblock] = " "
						layoutblock++
					}
				} else {
					for a := 0; a < next; a++ {
						horizvert()
						if layoutblockh > layouth/2 {
							break
						}
						layoutmap[layoutblock] = " "
						layoutblock += layoutw
					}
				}
			} else {
				choose := rInt(1, 4)
				if choose == 1 {
					for a := 0; a < next; a++ {
						horizvert()
						if layoutblockh < trackmaxl {
							break
						}
						layoutmap[layoutblock] = " "
						layoutblock -= layoutw
					}
				} else if choose == 2 {
					for a := 0; a < next; a++ {
						horizvert()
						if layoutblockv > layoutw-trackmaxl {
							break
						}
						layoutmap[layoutblock] = " "
						layoutblock++
					}
				} else if choose == 3 {
					for a := 0; a < next; a++ {
						horizvert()
						if layoutblockh > layouth/2 {
							break
						}
						layoutmap[layoutblock] = " "
						layoutblock += layoutw
					}
				}
			}

		} else {
			break
		}

	}

	/* draw square track
	for {
		layoutblock++
		horizvert()
		if layoutblockv < layoutw-minb {
			layoutmap[layoutblock] = " "
		} else {
			layoutmap[layoutblock] = " "
			break
		}
	}
	for {
		layoutblock += layoutw
		horizvert()
		if layoutblockh < layouth-minb {
			layoutmap[layoutblock] = " "
		} else {
			layoutmap[layoutblock] = " "
			break
		}
	}
	for {
		layoutblock--
		horizvert()
		if layoutblockv > minb {
			layoutmap[layoutblock] = " "
		} else {
			layoutmap[layoutblock] = " "
			break
		}
	}
	for {
		layoutblock -= layoutw
		horizvert()
		if layoutblockh > minb {
			layoutmap[layoutblock] = " "
		} else {
			break
		}
	}
	*/
	createdrawmap()
	createextras()
}
func createdrawmap() { // MARK: createlevel()
	for a := 0; a < drawmapa; a++ {
		drawmap[a] = "."
	}
	count := 0
	countlayout := 0
	for a := 0; a < layouta; a++ {
		if layoutmap[a] == " " {
			countscale := 0
			block := count
			for b := 0; b < scalea; b++ {
				drawmap[block] = " "
				block++
				countscale++
				if countscale == scale {
					countscale = 0
					block += drawmapw - scale
				}
			}
		}
		count += 5
		countlayout++
		if countlayout == layoutw {
			countlayout = 0
			count += drawmapw * 4
		}
	}

	tracktiles()
}
func input() { // MARK: input()

	if rl.IsKeyDown(rl.KeyRight) {
		playerdir = 2
		horizvert()
		if playerv < drawmapw-(trackminl*scale) {
			if drawmap[player] == " " {
				player++
			} else if drawmap[player] == "." {
				if framecount%2 == 0 {
					player++
				}
			}
		}
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		playerdir = 4
		horizvert()
		if playerv > trackminl*scale {
			if drawmap[player] == " " {
				player--
			} else if drawmap[player] == "." {
				if framecount%2 == 0 {
					player--
				}
			}
		}
	}
	if rl.IsKeyDown(rl.KeyDown) {
		if playerdir == 1 || playerdir == 2 {
			playerdir = 3
		} else if playerdir == 4 {
			playerdir = 5
		}
		horizvert()
		if playerh < drawmaph-(trackminl*scale) {
			player += drawmapw
		}
	}
	if rl.IsKeyDown(rl.KeyUp) {
		if playerdir == 2 || playerdir == 3 {
			playerdir = 1
		} else if playerdir == 4 || playerdir == 5 {
			playerdir = 6
		}
		horizvert()
		if playerh > trackminl*scale {
			player -= drawmapw
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else {
			camera.Zoom = 1.0
		}
	}

	if rl.IsKeyDown(rl.KeyKp8) {
		horizvert()
		if drawblocknexth > 0 {
			drawblocknext -= drawmapw
		}
	}
	if rl.IsKeyDown(rl.KeyKp2) {
		horizvert()
		if drawblocknexth < drawmaph-screenh {
			drawblocknext += drawmapw
		}
	}
	if rl.IsKeyDown(rl.KeyKp4) {
		horizvert()
		if drawblocknextv > 0 {
			drawblocknext--
		}
	}
	if rl.IsKeyDown(rl.KeyKp6) {
		horizvert()
		if drawblocknextv < drawmapw-screenw {
			drawblocknext++
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}
}
func setscreen() { // MARK: setscreen()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	screenw = (monitorw / 16) + 1
	screenh = (monitorh / 16) + 1
	screena = screenw * screenh

	camera.Zoom = 2.0
	camera.Target.X = 0.0
	camera.Target.Y = 0.0

}
func horizvert() { // MARK: horizvert()
	layoutblockh, layoutblockv = layoutblock/layoutw, layoutblock%layoutw
	drawblocknexth, drawblocknextv = drawblocknext/drawmapw, drawblocknext%drawmapw
	playerh, playerv = player/drawmapw, player%drawmapw
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Blue, 0.7))
	rl.DrawFPS(monw32-290, monh32-100)

	screenhTEXT := strconv.Itoa(screenh)
	screenwTEXT := strconv.Itoa(screenw)
	drawmapwTEXT := strconv.Itoa(drawmapw)
	drawmaphTEXT := strconv.Itoa(drawmaph)
	cameratargetXTEXT := fmt.Sprintf("%.0f", camera.Target.X)
	playervTEXT := strconv.Itoa(playerv)
	playerhTEXT := strconv.Itoa(playerh)

	rl.DrawText(screenwTEXT, monw32-290, 10, 10, rl.White)
	rl.DrawText("screenw", monw32-200, 10, 10, rl.White)
	rl.DrawText(screenhTEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("screenh", monw32-200, 20, 10, rl.White)
	rl.DrawText(drawmapwTEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("drawmapw", monw32-200, 30, 10, rl.White)
	rl.DrawText(drawmaphTEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("drawmaph", monw32-200, 40, 10, rl.White)
	rl.DrawText(cameratargetXTEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("cameratargetX", monw32-200, 50, 10, rl.White)
	rl.DrawText(playervTEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("playerv", monw32-200, 60, 10, rl.White)
	rl.DrawText(playerhTEXT, monw32-290, 70, 10, rl.White)
	rl.DrawText("playerh", monw32-200, 70, 10, rl.White)

}
func pconsole() { // MARK: pconsole()
	count := 0
	for a := 0; a < layouta; a++ {
		print(layoutmap[a])
		count++
		if count == layoutw {
			count = 0
			println()
		}
	}
	println()
	print("maindir ")
	print(maindir)
	println()
	print("dir ")
	print(dir)
	println()
	print("startblockv ")
	print(startblockv)
	println()
	print("startblockh ")
	print(startblockh)
	println()
	print("layoutblockh ")
	print(layoutblockh)
	println()
	print("layoutblockv ")
	print(layoutblockv)
	println()
	print("drawmapa ")
	print(len(drawmap))
	println()
	print("player ")
	print(player)

}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
