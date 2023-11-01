package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	win := myApp.NewWindow("sample")
	win.Resize(fyne.NewSize(400, 300))
	win.ShowAndRun()
}
