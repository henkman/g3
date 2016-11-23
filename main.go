package main

import (
	"log"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_mixer"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func main() {
	runtime.LockOSThread()
	log := log.New(os.Stderr, "", log.Ltime|log.Lmicroseconds)
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatal(err)
	}
	if err := ttf.Init(); err != nil {
		log.Fatal(err)
	}
	if err := mix.Init(0); err != nil {
		log.Fatal(err)
	}
	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatal(err)
	}
	win, err := sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		1024, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	render, err := sdl.CreateRenderer(
		win, -1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatal(err)
	}
	gd := GameData{Log: log, Win: win, Render: render}
	var game Game
	game.Init(gd)
	var scene Scene = &game
	for scene != nil {
		scene = scene.Run(gd)
	}
	render.Destroy()
	win.Destroy()
	mix.Quit()
	ttf.Quit()
	sdl.Quit()
}
