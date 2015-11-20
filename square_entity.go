package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/xamino/orbengine"
)

/*
square is a simple test entity.
*/
type square struct {
	id    string
	rect  *sdl.Rect
	color *sdl.Color
	pos   *sdl.Point
	rot   float64
}

func newSquare(id string, posX, posY, width, height int32, r, g, b uint8, rotation float64) *square {
	pos := &sdl.Point{X: posX, Y: posY}
	rect := &sdl.Rect{
		X: posX - (width / 2),
		Y: posY - (height / 2),
		W: width,
		H: height}
	color := &sdl.Color{R: r, G: g, B: b, A: 255}
	return &square{
		id:    id,
		rect:  rect,
		color: color,
		pos:   pos,
		rot:   rotation}
}

func (s *square) Identification() string {
	return s.id
}

func (s *square) Render(renderer *orbengine.Renderer) {
	renderer.SetDrawColor(s.color.R, s.color.G, s.color.B, s.color.A)
	renderer.FillRect(nil) // fill nil --> fill everything
}

func (s *square) Redraw() bool {
	// green needs to be redrawn frequently because it is color animated
	if s.color.G == 255 {
		return true
	}
	return false
}

func (s *square) Position() *sdl.Point {
	return s.pos
}

func (s *square) Offset() *sdl.Point {
	return nil
}

func (s *square) Scale() int {
	return 1
}

func (s *square) Width() int {
	return int(s.rect.W)
}

func (s *square) Height() int {
	return int(s.rect.H)
}

func (s *square) Rotation() float64 {
	return s.rot
}

func (s *square) Layer() int {
	return -1
}

func (s *square) Action() {
	// have red move (green check to avoid green moving every now and then)
	if s.color.R == 255 && s.color.G == 0 {
		// reset if reached border
		if s.pos.X < int32(-s.Width()) {
			// log.Println("Red: reset")
			s.pos.X = 600
			s.rect.X = 600
		}
		// TODO couple these two values somehow
		s.pos.X -= 5
		s.rect.X -= 5
	}
	// have blue continously turn
	if s.color.B == 255 {
		if s.rot >= 360 {
			// log.Println("Blue: reset")
			s.rot = 0
		}
		s.rot = (s.rot + 5)
	}
	// have green animate color
	if s.color.G == 255 {
		if s.color.R >= 255 {
			// log.Println("Green: reset")
			s.color.R = 0
		} else {
			s.color.R += 5
		}
	}
}
