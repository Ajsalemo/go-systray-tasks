package main

import (
	"bytes"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
)

func invokeHotKeys() {
	pasteBacklogTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyS)
	pasteBacklogBody := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	pasteAgedTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyD)
	pasteAgedBody := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyD)
	// Register backlog title hotkey
	errBacklogTitle := pasteBacklogTitle.Register()
	if errBacklogTitle != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + S")
		zap.L().Fatal(errBacklogTitle.Error())
		return
	}
	// Register backlog body hotkey
	errBacklogBody := pasteBacklogBody.Register()
	if errBacklogBody != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + SHIFT + S")
		zap.L().Fatal(errBacklogBody.Error())
		return
	}
	// Register aged title hotkey
	errAgedTitle := pasteAgedTitle.Register()
	if errAgedTitle != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + D")
		zap.L().Fatal(errAgedTitle.Error())
		return
	}
	// Register aged body hotkey
	errAgedBody := pasteAgedBody.Register()
	if errAgedBody != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + SHIFT + D")
		zap.L().Fatal(errAgedBody.Error())
		return
	} // Run this on a infinite loop to watch for key events
	// Otherwise this doesn't watch further key events aside from the first one
	for {
		select {
		case <-pasteBacklogTitle.Keydown():
			zap.L().Info("hotkey: CTRL + S is down")
		// Create a new backlog title and paste it to the clipboard
		case <-pasteBacklogTitle.Keyup():
			backlogTitle := os.Getenv("BACKLOG_TITLE_PREFIX") + " | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(backlogTitle).Bytes())
			zap.L().Info("hotkey: CTRL + S is up")
		case <-pasteBacklogBody.Keydown():
			zap.L().Info("hotkey: CTRL + SHIFT + S is down")
		// Read the file and paste it to the clipboard for the backlog body
		case <-pasteBacklogBody.Keyup():
			backlogTitle := os.Getenv("BACKLOG_BODY_FILE_PATH")
			backlogBodyContent, err := os.ReadFile(backlogTitle)
			if err != nil {
				zap.L().Error("failed to read file for BACKLOG_BODY_FILE_PATH")
				zap.L().Error("is the filesytem read-only or does the file exist?")
				zap.L().Fatal(err.Error())
				return
			}

			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(string(backlogBodyContent)).Bytes())
			zap.L().Info("hotkey: CTRL + SHIFT + S is up")
		case <-pasteAgedTitle.Keydown():
			zap.L().Info("hotkey: CTRL + D is down")
		// Create a new backlog title and paste it to the clipboard
		case <-pasteAgedTitle.Keyup():
			agedTitle := os.Getenv("AGED_TITLE_PREFIX") + " | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(agedTitle).Bytes())
			zap.L().Info("hotkey: CTRL + D is up")
		case <-pasteAgedBody.Keydown():
			zap.L().Info("hotkey: CTRL + SHIFT + D is down")
		// Read the file and paste it to the clipboard for the aged body
		case <-pasteAgedBody.Keyup():
			agedTitle := os.Getenv("AGED_BODY_FILE_PATH")
			agedBodyContent, err := os.ReadFile(agedTitle)
			if err != nil {
				zap.L().Error("failed to read file for AGED_BODY_FILE_PATH")
				zap.L().Error("is the filesytem read-only or does the file exist?")
				zap.L().Fatal(err.Error())
				return
			}

			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(string(agedBodyContent)).Bytes())
			zap.L().Info("hotkey: CTRL + SHIFT + D is up")
		}
	}
}

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	app := fiber.New()
	// TODO - expose pprof endpoints for debugging/metrics to see if these for loops are going to cause a problem
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Run these functions in a goroutine
	// These keybind watchers are being watched on a infinite for loop. It seems better to run them in a goroutine because of this
	go func() {
		invokeHotKeys()
	}()

	fiberErr := app.Listen(":3000")

	if fiberErr != nil {
		zap.L().Fatal(fiberErr.Error())
	}
}
