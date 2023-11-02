package gui

import (
	"goids/goids"
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func Run(width, height int, n int, speed float64, force float64, sight float64) {
	a := app.New()
	w := a.NewWindow("Goids")
	w.Resize(fyne.NewSize(float32(width), float32(height)))
	w.CenterOnScreen()

	e := goids.CreateEnv(float64(width), float64(height), n, speed, force, sight)

	go func() {
		img := e.RenderImage()
		imageWidget := canvas.NewImageFromImage(img)
		w.SetContent(imageWidget)
		for range time.Tick(time.Second / 60) {
			tick(&e, imageWidget)
		}
	}()

	w.ShowAndRun()
}

func tick(e *goids.Environment, imageWidget *canvas.Image) {
	e.Update()
	updateImage(imageWidget, e.RenderImage())
}

func updateImage(imageWidget *canvas.Image, img image.Image) {
	imageWidget.Image = img
	imageWidget.Refresh()
}
