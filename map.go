package main

import "github.com/veandco/go-sdl2/sdl"

type Collision uint8

const (
	Collision_No Collision = iota
	Collision_Yes
	Collision_OnlyDown
)

type Tile uint

type Map struct {
	Name         string
	Width        uint
	Height       uint
	Tile         []Tile
	Collision    []Collision
	Tilemap      *sdl.Texture
	TilemapWidth uint
	Spawn        struct {
		X, Y uint
	}
}

func (m *Map) Render(render *sdl.Renderer) {
	var src, dst sdl.Rect
	src.W = 32
	src.H = 32
	dst.W = 32
	dst.H = 32
	var x, y uint
	for y = 0; y < m.Height; y++ {
		for x = 0; x < m.Width; x++ {
			dst.X = int32(x) * dst.W
			dst.Y = int32(y) * dst.H
			t := m.Tile[y*m.Width+x]
			t--
			src.X = int32(uint(t)%m.TilemapWidth) * 32
			src.Y = int32(uint(t)/m.TilemapWidth) * 32
			render.Copy(m.Tilemap, &src, &dst)
		}
	}
}
