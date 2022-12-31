package server

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"

	"github.com/librespot-org/librespot-golang/librespot"
	"github.com/librespot-org/librespot-golang/librespot/core"
	"github.com/librespot-org/librespot-golang/librespot/player"
	
)

var (
	clientId =  "e76ef41e470a45cb923f6f77aaae5bf1"
	clientSecret = "f170d1874c8d46d3911f6066b6ac822a"
)

func StartServer() {
	session := make(chan core.Session)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		urlPath := "https://accounts.spotify.com/authorize?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&redirect_uri=http://localhost:8888/callback" +
		"&scope=streaming"
		openbrowser(urlPath)
		go func() {
			sess, _ := librespot.LoginOAuth("device", clientId, clientSecret)
			session <- *sess
		}()
	})

	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		sess := <-session
		w.Write([]byte(sess.DeviceId()))
		uri := strings.TrimPrefix(r.URL.Path, "/convert/")
		track, _ := sess.Mercury().GetTrack(uri)
		if track == nil {
			alt := track.GetAlternative()
			track = alt[0]
		}
		file := track.GetFile()
		file[0].

		audioFile, _ := pl

	})

	// http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
	// 	params := r.URL.Query()
	// 	auth, err := core.GetOauthAccessToken(params.Get("code"), "http://localhost:3001/callback", clientId, clientSecret)
	// 	if err != nil {
	// 		fmt.Fprintf(w, "Error getting token %q", err)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "Got token, loggin in")
	// 	ch <- *auth
	// })


	log.Print("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}


func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
