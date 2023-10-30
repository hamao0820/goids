package main

import (
	"math/rand"
)

func main() {
	goids := make([]Goid, 30)
	width := 640.0
	height := 480.0
	for i := range goids {
		position := CreateVector(rand.Float64()*width, rand.Float64()*height)
		velocity := CreateVector(rand.Float64()*2-1, rand.Float64()*2-1)
		velocity.Scale(rand.Float64()*4 - rand.Float64()*2)
		maxSpeed := 4
		maxForce := 2
		goids[i] = Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce)}
	}

	for i := 0; i < 1000; i++ {
		for i := 0; i < len(goids); i++ {
			goid := &goids[i]
			goid.Flock(goids)
			goid.Update(width, height)
		}
	}
}
