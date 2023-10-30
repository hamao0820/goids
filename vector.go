package main

type Vector struct {
	X int
	Y int
}

func CreateVector(x, y int) Vector {
	return Vector{X: x, Y: y}
}
