package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var ypos int32

func main() {
	log.Println("Hello world!")

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				log.Println("quiting")
				running = false
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_UP:
					log.Println("UP pressed!")
					ypos -= 10
				case sdl.K_DOWN:
					log.Println("DOWN pressed!")
					ypos += 10
				}
			}
		}
		render(surface)
		window.UpdateSurface()
	}

	sdl.Quit()
}

func render(surface *sdl.Surface) {
	// TODO redrawing the entire surface is ugly. Look into how SDL allows intelligent updating
	surface.FillRect(&sdl.Rect{0, 0, surface.W, surface.H}, 0x00000000)
	surface.FillRect(&sdl.Rect{0, 0, 200, 200}, 0xffff0000)
	surface.FillRect(&sdl.Rect{200, ypos, 200, 200}, 0x0000ffff)
}
