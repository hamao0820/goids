package main

import (
	"goids/goids"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	width, height := 640, 480

	myApp := app.New()
	window := myApp.NewWindow("Image")
	window.Resize(fyne.NewSize(float32(width), float32(height)))

	e := goids.CreateEnv(float64(width), float64(height), 30, 4, 2)

	go func() {
		img := e.RenderImage()
		imageWidget := canvas.NewImageFromImage(img)
		window.SetContent(imageWidget)
	}()

	window.ShowAndRun()
}
