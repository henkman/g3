package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_mixer"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	RES_DIR = "res/"
	IMG_DIR = RES_DIR + "img/"
	SND_DIR = RES_DIR + "snd/"
	FNT_DIR = RES_DIR + "fnt/"
)

// TODO cache textures
func LoadTexture(gd GameData, res string) *sdl.Texture {
	gd.Log.Println("loading texture", res)
	tex, err := img.LoadTexture(gd.Render, IMG_DIR+res)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return tex
}

func LoadSound(gd GameData, res string) *mix.Chunk {
	gd.Log.Println("loading sound", res)
	snd, err := mix.LoadWAV(SND_DIR + res)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return snd
}

func LoadFont(gd GameData, res string, size int) *ttf.Font {
	gd.Log.Println("loading font", res)
	fnt, err := ttf.OpenFont(FNT_DIR+res, size)
	if err != nil {
		gd.Log.Fatal(err)
	}
	return fnt
}
