package main

import (
	"goids/goids"
	"time"

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
		for {
			e.Update()
			imageWidget.Image = e.RenderImage()
			imageWidget.Refresh()

			time.Sleep(30 * time.Millisecond)
		}
	}()

	window.ShowAndRun()
}
