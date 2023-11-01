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

	a := app.New()
	w := a.NewWindow("Image")
	w.Resize(fyne.NewSize(float32(width), float32(height)))
	w.CenterOnScreen()

	e := goids.CreateEnv(float64(width), float64(height), 30, 3, 2)

	go func() {
		img := e.RenderImage()
		imageWidget := canvas.NewImageFromImage(img)
		w.SetContent(imageWidget)
		for {
			e.Update()
			imageWidget.Image = e.RenderImage()
			imageWidget.Refresh()

			time.Sleep(time.Second / 60)
		}
	}()

	w.ShowAndRun()
	a.Quit()
}
