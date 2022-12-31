package ui

import (
	"github.com/fyne-io/fyne/v2/app"
	"github.com/fyne-io/fyne/v2/container"
	"github.com/fyne-io/fyne/v2/widget"
)

func LoadApp() {
	a := app.New()
	w := a.NewWindow("Hello")
	
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}