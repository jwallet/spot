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


	// fs := http.FileServer(http.Dir("./frontend/web-build"))
	// http.Handle("/", fs)

	// openbrowser("http://localhost:3000")
	
	// log.Print("Listening on :3000...")
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// func openbrowser(url string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }