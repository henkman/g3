package main

import "github.com/veandco/go-sdl2/sdl"

type Player struct {
	Position      Vec
	Velocity      Vec
	Height        int32
	Width         int32
	CanDoubleJump bool
}

func (p *Player) Render(render *sdl.Renderer) {
	render.SetDrawColor(0x00, 0xFF, 0x00, 0xFF)
	render.FillRect(&sdl.Rect{int32(p.Position.X), int32(p.Position.Y), p.Width, p.Height})
}

type Game struct {
	Player Player
	Map    Map
}

func (g *Game) Init(gd GameData) {
	p := &g.Player
	p.Width = TileSize
	p.Height = TileSize * 2
	g.Map = LoadMap(gd, "1.json")
	p.Position.X = float32(g.Map.Spawn.X)
	p.Position.Y = float32(g.Map.Spawn.Y)
}

func (g *Game) Run(gd GameData) Scene {
	for {
		var kp struct {
			JumpUp bool
		}
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch t := e.(type) {
			case *sdl.QuitEvent:
				return nil
			case *sdl.KeyUpEvent:
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_ESCAPE:
					return nil
				case sdl.SCANCODE_SPACE:
					kp.JumpUp = true
				}
			}
		}
		{
			const GRAVITY = 1
			w, h, _ := gd.Render.GetRendererOutputSize()
			p := &g.Player

			kb := sdl.GetKeyboardState()
			if kb[sdl.SCANCODE_LEFT] != 0 {
				p.Velocity.X += -5
			} else if kb[sdl.SCANCODE_RIGHT] != 0 {
				p.Velocity.X += 5
			}
			if kp.JumpUp {
				p.Velocity.Y -= 20
			}
			p.Velocity.Y += GRAVITY
			p.Velocity.X = Clamp(p.Velocity.X, -5, 5)
			p.Velocity.Y = Clamp(p.Velocity.Y, -40, 40)
			p.Position = p.Position.Add(p.Velocity)

			p.Position.X = Clamp(p.Position.X, 0, float32(w-int(p.Width)))
			p.Position.Y = Clamp(p.Position.Y, 0, float32(h-int(p.Height)))

			for {
				c := g.Map.Collides(uint32(p.Position.X), uint32(p.Position.Y),
					uint32(p.Width), uint32(p.Height))
				if c.Floor {
					p.Position.Y = ((p.Position.Y - 1) / TileSize) * TileSize
					p.Velocity.Y = 0
					p.Velocity.X = 0
				} else {
					break
				}
			}

		}
		gd.Render.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
		gd.Render.Clear()
		{
			g.Map.Render(gd.Render)
			g.Player.Render(gd.Render)
		}
		gd.Render.Present()
	}
}
