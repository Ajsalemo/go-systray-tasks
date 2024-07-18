package main

import (
	"log"

	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	mainthread.Init(fn)
}

func fn() {
	pasteBacklogTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	err := pasteBacklogTitle.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}

	log.Printf("hotkey: %v is registered\n", pasteBacklogTitle)
	<-pasteBacklogTitle.Keydown()
	log.Printf("hotkey: %v is down\n", pasteBacklogTitle)

	if pasteBacklogTitle.Keydown() != nil {
		log.Printf("hotkey: %v is down again \n", pasteBacklogTitle)
	}

	if pasteBacklogTitle.Keyup() != nil {
		clipboard.Write(clipboard.FmtText, []byte("text data"))
		log.Printf("hotkey: %v is up\n", pasteBacklogTitle)
	}
}
