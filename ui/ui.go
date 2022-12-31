package ui

import (
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LoadApp(w fyne.Window) {
	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			res, _ := http.Get("http://localhost:3001/login")
			if res.StatusCode == 200 {
				hello.SetText("Logged in")
			}
		}),
	))
}