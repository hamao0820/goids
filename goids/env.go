package goids

import (
	"bytes"
	"embed"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math/rand"

	xdraw "golang.org/x/image/draw"
)

//go:embed img/*
var imgs embed.FS

type Environment struct {
	width    float64
	height   float64
	goidsNum int
	goids    []Goid
	maxSpeed float64
	maxForce float64

	frontImage image.Image
	SideImage  image.Image
	PinkImage  image.Image
}

func CreateEnv(width, height float64, n int, maxSpeed, maxForce float64, sight float64) Environment {
	goids := make([]Goid, n)
	for i := range goids {
		position := CreateVector(rand.Float64()*width, rand.Float64()*height)
		velocity := CreateVector(rand.Float64()*2-1, rand.Float64()*2-1)
		velocity.Scale(rand.Float64()*4 - rand.Float64()*2)
		maxSpeed := maxSpeed
		maxForce := maxForce

		var t ImageType

		r := rand.Float64()

		if r < 0.001 { // 0.1%
			t = Pink
		} else if r < 0.011 { // 1%
			t = Side
		} else {
			t = Front
		}

		goids[i] = Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce), sight: sight, imageType: t}
	}

	imgFront := loadImage("img/gopher-front.png")
	imgDstFront := image.NewRGBA(image.Rect(0, 0, int(float64(imgFront.Bounds().Dx())*32.0/float64(imgFront.Bounds().Dy())), 32)) // 高さを32に固定
	xdraw.CatmullRom.Scale(imgDstFront, imgDstFront.Bounds(), imgFront, imgFront.Bounds(), draw.Over, nil)

	imgSide := loadImage("img/gopher-side.png")
	imgDstSide := image.NewRGBA(image.Rect(0, 0, int(float64(imgSide.Bounds().Dx())*32.0/float64(imgSide.Bounds().Dy())), 32)) // 高さを32に固定
	xdraw.CatmullRom.Scale(imgDstSide, imgDstSide.Bounds(), imgSide, imgSide.Bounds(), draw.Over, nil)

	imgPink := loadImage("img/gopher-pink.png")
	imgDstPink := image.NewRGBA(image.Rect(0, 0, int(float64(imgPink.Bounds().Dx())*32.0/float64(imgPink.Bounds().Dy())), 32)) // 高さを32に固定
	xdraw.CatmullRom.Scale(imgDstPink, imgDstPink.Bounds(), imgPink, imgPink.Bounds(), draw.Over, nil)

	return Environment{width: width, height: height, goidsNum: n, goids: goids, maxSpeed: maxSpeed, maxForce: maxForce, frontImage: imgDstFront.SubImage(imgDstFront.Rect), SideImage: imgDstSide.SubImage(imgDstSide.Rect), PinkImage: imgDstPink.SubImage(imgDstPink.Rect)}
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

func (e Environment) RenderImage() image.Image {
	dest := image.NewRGBA(image.Rect(0, 0, int(e.Width()), int(e.Height())))
	for _, goid := range e.goids {
		var img image.Image
		switch goid.imageType {
		case Front:
			img = e.frontImage
		case Side:
			img = e.SideImage
		case Pink:
			img = e.PinkImage
		}

		p := image.Point{int(goid.position.X), int(goid.position.Y)}
		rectAngle := image.Rectangle{p.Sub(img.Bounds().Size().Div(2)), p.Add(img.Bounds().Size().Div(2))}
		draw.Draw(dest, rectAngle, img, image.Point{0, 0}, draw.Over)
	}
	return dest.SubImage(dest.Rect)
}

func (e Environment) Render() string {
	img := e.RenderImage()
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return fmt.Sprintf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}

func loadImage(path string) image.Image {
	f, err := imgs.ReadFile(path)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		panic(err)
	}
	return img
}
