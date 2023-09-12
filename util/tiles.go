package util

import rl "github.com/gen2brain/raylib-go/raylib"

//func int() {
//	rl.DrawRectangle(100, 100, 200, 150, rl.Blue)
//}

type Tile struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
}

func NewTile(pos rl.Vector2, size rl.Vector2, color rl.Color) Tile {
	return Tile{
		Position: pos,
		Size:     size,
		Color:    color,
	}
}

func (tile *Tile) Draw() {
	rl.DrawRectangleV(tile.Position, tile.Size, tile.Color)
}
