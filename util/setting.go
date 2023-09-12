package util

var Tile_size int
var Screen_width int32
var Screen_heigh int32

var LevelMap []string

func init() {
	Screen_width = 1200
	Screen_heigh = 700
	Tile_size = 64

	LevelMap = []string{
		"                            ",
		"                            ",
		"                            ",
		" XX    XXX            XX    ",
		" XX P                       ",
		" XXXX         XX         XX ",
		" XXXX       XX              ",
		" XX    X  XXXX    XX  XX    ",
		"       X  XXXX    XX  XXX   ",
		"    XXXX  XXXXXX  XX  XXXX  ",
		"XXXXXXXX  XXXXXX  XX  XXXX  ",
	}

}
