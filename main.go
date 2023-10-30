package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"

	"github.com/llgcode/draw2d/draw2dimg"
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
	clearScreen()
	hideCursor()

	for i := 0; i < 1000; i++ {
		for i := 0; i < len(goids); i++ {
			goid := &goids[i]
			goid.Flock(goids)
			goid.Update(width, height)
		}
		frame := draw(goids)
		fmt.Printf("loop: %d\n", i)
		printImage(frame.SubImage(frame.Rect))
	}
	showCursor()
}

// draw the goids
func draw(goids []Goid) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, 640, 480))
	gc := draw2dimg.NewGraphicContext(dest)
	for _, goid := range goids {
		gc.SetFillColor(color.RGBA{200, 200, 100, 255})
		gc.MoveTo(float64(goid.position.X), float64(goid.position.Y))
		gc.ArcTo(float64(goid.position.X), float64(goid.position.Y), float64(5), float64(5), 0, -math.Pi*2)
		// gc.LineTo(float64(goid.X-goid.Vx), float64(goid.Y-goid.Vy))
		gc.Close()
		gc.Fill()
	}
	return dest
}

// ANSI escape sequence codes to perform action on terminal
func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\x1b[?25h\n")
}

func clearScreen() {
	fmt.Print("\x1b[2J")
}

// this only works for iTerm!
func printImage(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a", imgBase64Str)
}
