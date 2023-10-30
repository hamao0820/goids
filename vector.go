package main

type Vector struct {
	X float64
	Y float64
}

func CreateVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}
