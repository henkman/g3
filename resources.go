package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_mixer"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	RES_DIR = "res/"
	IMG_DIR = "img/"
	SND_DIR = "snd/"
	FNT_DIR = "fnt/"
	MAP_DIR = "map/"
)

// TODO cache textures
func LoadTexture(gd GameData, res string) *sdl.Texture {
	gd.Log.Println("loading texture", res)
	return loadTexture(gd, filepath.Join(RES_DIR, IMG_DIR, res))
}

func loadTexture(gd GameData, file string) *sdl.Texture {
	tex, err := img.LoadTexture(gd.Render, file)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return tex
}

func LoadSound(gd GameData, res string) *mix.Chunk {
	gd.Log.Println("loading sound", res)
	snd, err := mix.LoadWAV(filepath.Join(RES_DIR, SND_DIR, res))
	if err != nil {
		gd.Log.Fatal(err)
	}
	return snd
}

func LoadFont(gd GameData, res string, size int) *ttf.Font {
	gd.Log.Println("loading font", res)
	fnt, err := ttf.OpenFont(filepath.Join(RES_DIR, FNT_DIR, res), size)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return fnt
}

func LoadMap(gd GameData, res string) Map {
	gd.Log.Println("loading map", res)
	fd, err := os.Open(filepath.Join(RES_DIR, MAP_DIR, res))
	if err != nil {
		gd.Log.Fatal(err)
	}
	var md struct {
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
		Height uint32 `json:"height"`
		Width  uint32 `json:"width"`
		Layers []struct {
			Data    []Tile `json:"data,omitempty"`
			Name    string `json:"name"`
			Objects []struct {
				Name     string `json:"name"`
				Rotation int    `json:"rotation"`
				X        uint32 `json:"x"`
				Y        uint32 `json:"y"`
			} `json:"objects,omitempty"`
		} `json:"layers"`
	}
	if err := json.NewDecoder(fd).Decode(&md); err != nil {
		fd.Close()
		gd.Log.Fatal(err)
	}
	fd.Close()
	var m Map
	m.Tilemap = loadTexture(gd, filepath.Join(RES_DIR, MAP_DIR, res[:strings.LastIndex(res, ".")]+".png"))
	{
		_, _, w, _, _ := m.Tilemap.Query()
		m.TilemapWidth = uint(w) / TileSize
	}
	m.Name = md.Properties.Name
	m.Width = md.Width
	m.Height = md.Height
	m.Tile = make([]Tile, m.Height*m.Width)
	m.Collision = make([]Collision, 0, m.Height*m.Width)
	for _, layer := range md.Layers {
		if layer.Name == "collision" {
			for _, col := range layer.Data {
				var c Collision
				if col == 0 {
					c = Collision_No
				} else {
					c = Collision_Yes
				}
				m.Collision = append(m.Collision, c)
			}
		} else if layer.Name == "objects" {
			for _, obj := range layer.Objects {
				if obj.Name == "spawn" {
					m.Spawn.X = obj.X
					m.Spawn.Y = obj.Y
				}
			}
		} else if layer.Name == "tiles" {
			copy(m.Tile, layer.Data)
		}
	}
	return m
}
