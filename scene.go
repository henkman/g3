package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type GameData struct {
	Log    *log.Logger
	Win    *sdl.Window
	Render *sdl.Renderer
}

type Scene interface {
	Init(GameData)
	Run(GameData) Scene
}
