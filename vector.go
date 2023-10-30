package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func CreateVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) ScalarMul(c float64) {
	v.X *= c
	v.Y *= c
}

func (v *Vector) Scale(l float64) {
	v.ScalarMul(l / v.Len())
}
