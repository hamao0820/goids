package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
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

	f, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		for i := 0; i < len(goids); i++ {
			goid := &goids[i]
			goid.Flock(goids)
			goid.Update(width, height)
		}
		frame := drawGopher(goids, img)
		fmt.Printf("loop: %d\n", i)
		printImage(frame.SubImage(frame.Rect))
	}
	showCursor()
}

// draw the goids
func drawGopher(goids []Goid, img image.Image) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for _, goid := range goids {
		p := image.Point{int(goid.position.X), int(goid.position.Y)}
		rectAngle := image.Rectangle{p.Sub(img.Bounds().Size().Div(2)), p.Add(img.Bounds().Size().Div(2))}
		draw.Draw(dest, rectAngle, img, image.Point{0, 0}, draw.Over)
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
	fmt.Printf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}
