package main

import (
	settings "Nauka1/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(settings.Screen_width, settings.Screen_heigh, "Kanciak the Game")

	//	for _, row := range settings.LevelMap {
	//	fmt.Println(row)
	//	}

	//rl.SetTargetFPS(60)
	tiles := []settings.Tile{
		settings.NewTile(rl.NewVector2(100, 100), rl.NewVector2(50, 50), rl.Red),
		settings.NewTile(rl.NewVector2(200, 200), rl.NewVector2(50, 50), rl.Blue),
		settings.NewTile(rl.NewVector2(300, 300), rl.NewVector2(50, 50), rl.Green),
	}
	redTileIndex := 0
	blueTileIndex := 1

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyRight) {
			tiles[redTileIndex].Position.X += 5
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			tiles[redTileIndex].Position.X -= 5
		}

		if rl.IsKeyDown(rl.KeyUp) {
			tiles[redTileIndex].Position.Y -= 5
		}

		if rl.IsKeyDown(rl.KeyDown) {
			tiles[redTileIndex].Position.Y += 5
		}

		if rl.IsKeyDown(rl.KeyD) {
			tiles[blueTileIndex].Position.X += 5
		}

		if rl.IsKeyDown(rl.KeyA) {
			tiles[blueTileIndex].Position.X -= 5
		}

		if rl.IsKeyDown(rl.KeyW) {
			tiles[blueTileIndex].Position.Y -= 5
		}

		if rl.IsKeyDown(rl.KeyS) {
			tiles[blueTileIndex].Position.Y += 5
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawFPS(int32(settings.Screen_width)-90, int32(settings.Screen_heigh)-30)
		rl.DrawText("Kanciak", settings.Screen_width/2, settings.Screen_heigh/2, 30, rl.Gray)

		// Draw the tiles
		for _, tile := range tiles {
			tile.Draw()
		}
		//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
