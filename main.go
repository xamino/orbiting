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
	// build test entities
	red := newSquare(50, 50, 25, 25, 255, 0, 0, 0)
	blue := newSquare(250, 250, 100, 100, 0, 0, 255, 0)
	green := newSquare(400, 400, 20, 20, 0, 255, 0, 22.5)
	_ = controller.AddEntity("red", red)
	_ = controller.AddEntity("blue", blue)
	_ = controller.AddEntity("green", green)
	// build player
	player := newPlayer(350, 350, 50, 50, 127, 127, 127, controller)
	_ = controller.AddEntity("player", player)
	// start controller to run
	// TODO this should be handled in one elegant function WITHIN Controller including events!
	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-tick:
			controller.Iterate()
		}
	}
}
