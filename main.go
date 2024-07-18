package main

import (
	"bytes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
)

// TODO - clean this up to be more realistic. this is a POC
func fn() {
	pasteBacklogTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	err := pasteBacklogTitle.Register()
	if err != nil {
		log.Fatalf("hotkey: failed to register hotkey: %v", err)
		return
	}
	// Run this on a infinite loop to watch for key events
	// Otherwise this doesn't watch further key events aside from the first one
	for {
		select {
		case <-pasteBacklogTitle.Keydown():
			log.Printf("hotkey: %v is down\n", pasteBacklogTitle)
		case <-pasteBacklogTitle.Keyup():
			backlogTitle := "test | backlog | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(backlogTitle).Bytes())
			log.Printf("hotkey: %v is up\n", pasteBacklogTitle)
		}
	}
}

func main() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	// TODO - expose pprof endpoints for debugging/metrics to see if these for loops are going to cause a problem
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Run these functions in a goroutine
	// These keybind watchers are being watched on a infinite for loop. It seems better to run them in a goroutine because of this
	go func() {
		fn()
	}()

	log.Fatal(app.Listen(":3000"))
}
