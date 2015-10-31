package main

import (
	"log"
	"time"

	"github.com/xamino/orbengine"
)

// TODO camera as entity! makes for really simple render occlusion

func main() {
	// build controller
	controller, err := orbengine.Build()
	if err != nil {
		log.Println("Failed to start engine:", err)
		return
	}
	// destroy it when done
	defer controller.Destroy()
	// build two test entities
	red := newSquare(50, 50, 25, 25, 255, 0, 0)
	blue := newSquare(250, 250, 100, 100, 0, 0, 255)
	controller.AddEntity("red", red)
	controller.AddEntity("blue", blue)
	// TODO start controller to run
	timeout := time.Tick(5 * time.Second)
	tick := time.Tick(250 * time.Millisecond)
	for {
		select {
		case <-timeout:
			log.Println("Quit!")
			break
		case <-tick:
			controller.Iterate()
			log.Println("Drew")
		}
	}
}

/*
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
*/
