package main

import (
	"image"
	"image/draw"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	width, height := 100, 100

	myApp := app.New()
	window := myApp.NewWindow("Image")
	window.Resize(fyne.NewSize(float32(width), float32(height)))

	go func() {
		f, err := os.Open("test.png")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		dest := image.NewRGBA(image.Rect(0, 0, width, height))
		p := image.Point{rand.Intn(width), rand.Intn(height)}
		rectAngle := image.Rectangle{p.Sub(img.Bounds().Size().Div(2)), p.Add(img.Bounds().Size().Div(2))}
		draw.Draw(dest, rectAngle, img, image.Point{0, 0}, draw.Over)

		img_ := dest.SubImage(dest.Rect)
		imageWidget := canvas.NewImageFromImage(img_)
		imageWidget.Image = img_
		window.SetContent(imageWidget)
		for {
			dest := image.NewRGBA(image.Rect(0, 0, width, height))
			p := image.Point{rand.Intn(width), rand.Intn(height)}
			rectAngle := image.Rectangle{p.Sub(img.Bounds().Size().Div(2)), p.Add(img.Bounds().Size().Div(2))}
			draw.Draw(dest, rectAngle, img, image.Point{0, 0}, draw.Over)

			img_ := dest.SubImage(dest.Rect)
			imageWidget.Image = img_
			imageWidget.Refresh()

			time.Sleep(100 * time.Millisecond)
		}
	}()

	window.ShowAndRun()
}
