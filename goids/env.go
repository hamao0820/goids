package goids

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

type Environment struct {
	width    float64
	height   float64
	goidsNum int
	goids    []Goid
	maxSpeed float64
	maxForce float64

	image image.Image
}

func CreateEnv(width, height float64, n int, maxSpeed, maxForce float64) Environment {
	goids := make([]Goid, 30)
	for i := range goids {
		position := CreateVector(rand.Float64()*width, rand.Float64()*height)
		velocity := CreateVector(rand.Float64()*2-1, rand.Float64()*2-1)
		velocity.Scale(rand.Float64()*4 - rand.Float64()*2)
		maxSpeed := 4
		maxForce := 2
		goids[i] = Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce)}
	}

	f, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return Environment{width: width, height: height, goidsNum: n, goids: goids, maxSpeed: maxSpeed, maxForce: maxForce, image: img}
}

func (e *Environment) Update() {
	for i := 0; i < len(e.goids); i++ {
		goid := &e.goids[i]
		goid.Flock(e.goids)
		goid.Update(e.width, e.height)
	}
}

func (e Environment) Goids() []Goid {
	return e.goids
}

func (e Environment) GoidsNum() int {
	return e.goidsNum
}

func (e Environment) Width() float64 {
	return e.width
}

func (e Environment) Height() float64 {
	return e.height
}

func (e Environment) Render() string {
	dest := image.NewRGBA(image.Rect(0, 0, int(e.Width()), int(e.Height())))
	for _, goid := range e.goids {
		p := image.Point{int(goid.position.X), int(goid.position.Y)}
		rectAngle := image.Rectangle{p.Sub(e.image.Bounds().Size().Div(2)), p.Add(e.image.Bounds().Size().Div(2))}
		draw.Draw(dest, rectAngle, e.image, image.Point{0, 0}, draw.Over)
	}
	img := dest.SubImage(dest.Rect)
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return fmt.Sprintf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}
