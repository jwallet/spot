package main

import (
	"fmt"

	"github.com/librespot-org/librespot-golang/librespot/core"
	"github.com/librespot-org/librespot-golang/librespot/mercury"
	"github.com/librespot-org/librespot-golang/librespot/player"
)

type MobileSession struct {
	session *core.Session
	player  *MobilePlayer
	mercury *MobileMercury
}

type MobileMercury struct {
	mercury *mercury.Client
}

type MobilePlayer struct {
	player *player.Player
}

func login() (*MobileSession, error) {
	sess, err := core.LoginOAuth("deviceTest", "e76ef41e470a45cb923f6f77aaae5bf1", "f170d1874c8d46d3911f6066b6ac822a")
	if err != nil {
		return nil, err
	}

	fmt.Println(sess.Username())

	return initSessionImpl(sess)
}

func initSessionImpl(sess *core.Session) (*MobileSession, error) {
	return &MobileSession{
		session: sess,
		player:  createMobilePlayer(sess),
		mercury: createMobileMercury(sess),
	}, nil
}

func createMobilePlayer(session *core.Session) *MobilePlayer {
	return &MobilePlayer{
		player: session.Player(),
	}
}

func createMobileMercury(session *core.Session) *MobileMercury {
	return &MobileMercury{
		mercury: session.Mercury(),
	}
}
