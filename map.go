package main

import "github.com/veandco/go-sdl2/sdl"

type Collision uint8

const (
	Collision_No Collision = iota
	Collision_Yes
	Collision_OnlyDown
)

type CollisionDirection struct {
	Top   bool
	Floor bool
	Left  bool
	Right bool
}

const TileSize = 32

type Tile uint

type Map struct {
	Name         string
	Width        uint32
	Height       uint32
	Tile         []Tile
	Collision    []Collision
	Tilemap      *sdl.Texture
	TilemapWidth uint
	Spawn        struct {
		X, Y uint32
	}
}

func (m *Map) Collides(x, y, w, h uint32) CollisionDirection {
	x /= TileSize
	y /= TileSize
	w /= TileSize
	h /= TileSize

	return CollisionDirection{
		Top:   m.Collision[y*m.Width+x] == Collision_Yes,
		Floor: m.Collision[(y+h)*m.Width+x] == Collision_Yes,
		Left:  m.Collision[y*m.Width+x] == Collision_Yes,
		Right: m.Collision[y*m.Width+x+w] == Collision_Yes,
	}
}

func (m *Map) Render(render *sdl.Renderer) {
	var src, dst sdl.Rect
	src.W = TileSize
	src.H = TileSize
	dst.W = TileSize
	dst.H = TileSize
	var x, y uint32
	for y = 0; y < m.Height; y++ {
		for x = 0; x < m.Width; x++ {
			dst.X = int32(x) * dst.W
			dst.Y = int32(y) * dst.H
			t := m.Tile[y*m.Width+x]
			t--
			src.X = int32(uint(t)%m.TilemapWidth) * TileSize
			src.Y = int32(uint(t)/m.TilemapWidth) * TileSize
			render.Copy(m.Tilemap, &src, &dst)
		}
	}
}
