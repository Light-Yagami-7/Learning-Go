package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	aapp := app.New()
	appWindow := aapp.NewWindow("Network Manager")

	text := widget.NewLabel("JioFiber-5G")
	appWindow.SetContent(text)

	appWindow.ShowAndRun()
}

