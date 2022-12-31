package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/jwallet/spot/server"
	"github.com/jwallet/spot/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("Spot")
	w.Resize(fyne.NewSize(800, 600))

	go ui.LoadApp(w)
	w.Show()
	
	go server.StartServer()

	a.Run()
}