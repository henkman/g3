package main

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
}

func (g *Game) Init(gd GameData) {

}

func (g *Game) Run(gd GameData) Scene {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			if _, ok := e.(*sdl.QuitEvent); ok {
				return nil
			} else if kd, ok := e.(*sdl.KeyUpEvent); ok && kd.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
				return nil
			}
		}

		gd.Render.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
		gd.Render.Clear()
		gd.Render.Present()
	}
}
