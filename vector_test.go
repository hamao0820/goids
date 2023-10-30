package main

import (
	"math"
	"testing"
)

func TestLen(t *testing.T) {
	type args struct {
		v Vector
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test 1", args{Vector{3, 4}}, 5},
		{"Test 2", args{Vector{0, 0}}, 0},
		{"Test 3", args{Vector{3, -1}}, math.Sqrt(10)},
		{"Test 4", args{Vector{-1, 2}}, math.Sqrt(5)},
		{"Test 5", args{Vector{2, 2}}, math.Sqrt(8)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.v.Len(); got != tt.want {
				t.Errorf("Vector.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScalarMul(t *testing.T) {
	type args struct {
		v Vector
		c float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, 2}, Vector{6, 8}},
		{"Test 2", args{Vector{0, 0}, 2}, Vector{0, 0}},
		{"Test 3", args{Vector{3, -1}, 2}, Vector{6, -2}},
		{"Test 4", args{Vector{-1, 2}, 2}, Vector{-2, 4}},
		{"Test 5", args{Vector{2, 2}, 2}, Vector{4, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.ScalarMul(tt.args.c)
			if tt.args.v != tt.want {
				t.Errorf("Vector.ScalarMul() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		v Vector
		l float64
	}

	tests := []struct {
		name string
		args args
		want Vector
	}{
		{"Test 1", args{Vector{3, 4}, 2}, Vector{6.0 / 5.0, 8.0 / 5.0}},
		{"Test 2", args{Vector{0, 0}, 2}, Vector{0.0, 0.0}},
		{"Test 3", args{Vector{3, -1}, 2}, Vector{6 / math.Sqrt(10), -2 / math.Sqrt(10)}},
		{"Test 4", args{Vector{-1, 2}, 2}, Vector{-2 / math.Sqrt(5), 4 / math.Sqrt(5)}},
		{"Test 5", args{Vector{2, 2}, 2}, Vector{4 / math.Sqrt(8), 4 / math.Sqrt(8)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.v.Scale(tt.args.l)
			if (tt.args.v.X-tt.want.X) > 0.0000001 || (tt.args.v.Y-tt.want.Y) > 0.0000001 {
				t.Errorf("Vector.Scale() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
