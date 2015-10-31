package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

/*
square is a simple test entity.
*/
type square struct {
	rect  *sdl.Rect
	color *sdl.Color
	pos   *sdl.Point
}

func newSquare(posX, posY, width, height int32, r, g, b uint8) *square {
	pos := &sdl.Point{X: posX, Y: posY}
	rect := &sdl.Rect{
		X: posX - (width / 2),
		Y: posY - (height / 2),
		W: width,
		H: height}
	color := &sdl.Color{R: r, G: g, B: b, A: 255}
	return &square{
		rect:  rect,
		color: color,
		pos:   pos}
}

// Drawable interface
func (s *square) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(s.color.R, s.color.G, s.color.B, s.color.A)
	renderer.FillRect(s.rect)
}

// Boundable interface
func (s *square) Bounds() *sdl.Rect {
	return s.rect
}

// Action interface
func (s *square) Action() {
	//TODO update pos and update draw thingy?
	// FIXME: hack: only red square
	if s.color.R == 255 {
		// reset if reached border
		if s.pos.X < 0 {
			log.Println("reset!")
			s.pos.X = 100
			s.rect.X = 100
		}
		// TODO couple these two values somehow
		s.pos.X -= 10
		s.rect.X -= 10
	}
}
