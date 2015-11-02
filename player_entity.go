package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/xamino/orbengine"
)

/*
player is a simple controllable entity.
*/
type player struct {
	rect      *sdl.Rect
	color     *sdl.Color
	moveup    bool
	movedown  bool
	moveleft  bool
	moveright bool
}

func newPlayer(posX, posY, width, height int32, r, g, b uint8, c *orbengine.Controller) *player {
	rect := &sdl.Rect{
		X: posX - (width / 2),
		Y: posY - (height / 2),
		W: width,
		H: height}
	color := &sdl.Color{R: r, G: g, B: b, A: 255}
	p := &player{
		rect:  rect,
		color: color}
	// get key channels
	_ = c.RegisterKey("Up", func() { p.moveup = true }, func() { p.moveup = false })
	_ = c.RegisterKey("Down", func() { p.movedown = true }, func() { p.movedown = false })
	_ = c.RegisterKey("Left", func() { p.moveleft = true }, func() { p.moveleft = false })
	_ = c.RegisterKey("Right", func() { p.moveright = true }, func() { p.moveright = false })
	return p
}

func (p *player) Render(renderer *orbengine.Renderer) {
	renderer.SetDrawColor(p.color.R, p.color.G, p.color.B, p.color.A)
	renderer.FillRect(nil) // fill nil --> fill everything
}

func (p *player) Redraw() bool {
	return false
}

func (p *player) Position() *sdl.Point {
	return &sdl.Point{
		X: p.rect.X + (p.rect.W / 2),
		Y: p.rect.Y + (p.rect.H / 2)}
}

func (p *player) Offset() *sdl.Point {
	return nil
}

func (p *player) Scale() int {
	return 1
}

func (p *player) Width() int {
	return int(p.rect.W)
}

func (p *player) Height() int {
	return int(p.rect.H)
}

func (p *player) Rotation() float64 {
	return 0
}

func (p *player) Action() {
	var delta int32 = 10
	if p.moveup {
		p.rect.Y -= delta
	}
	if p.movedown {
		p.rect.Y += delta
	}
	if p.moveleft {
		p.rect.X -= delta
	}
	if p.moveright {
		p.rect.X += delta
	}
}
