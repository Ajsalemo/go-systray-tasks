package main

import (
	"bytes"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"

	controllers "gtpl/controllers"
)

func invokeHotKeys() {
	pasteBacklogTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyW)
	pasteBacklogBody := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyW)
	pasteAgedTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyE)
	pasteAgedBody := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyE)
	pasteFDRTitle := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyR)
	pasteFDRBody := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyR)
	
	// If DISABLE_BACKLOG_TITLE is "true" / "TRUE" or 1 - disable the hotkey
	// Otherwise, register backlog title hotkey
	if os.Getenv("DISABLE_BACKLOG_TITLE") == "true" || os.Getenv("DISABLE_BACKLOG_TITLE") == "TRUE" || os.Getenv("DISABLE_BACKLOG_TITLE") == "1" {
		zap.L().Info("hotkey: CTRL + W is disabled and unregistered")
	} else {
		// Register backlog title hotkey
		errBacklogTitle := pasteBacklogTitle.Register()
		if errBacklogTitle != nil {
			zap.L().Error("hotkey: failed to register hotkey: CTRL + W")
			zap.L().Fatal(errBacklogTitle.Error())
			return
		}
	}	
	// If DISABLE_BACKLOG_BODY is "true" / "TRUE" or 1 - disable the hotkey
	// Otherwise, register backlog body hotkey
	if os.Getenv("DISABLE_BACKLOG_BODY") == "true" || os.Getenv("DISABLE_BACKLOG_BODY") == "TRUE" || os.Getenv("DISABLE_BACKLOG_BODY") == "1" {
		zap.L().Info("hotkey: CTRL + SHIFT + W is disabled and unregistered")
	} else {
		// Register backlog body hotkey
		errBacklogBody := pasteBacklogBody.Register()
		if errBacklogBody != nil {
			zap.L().Error("hotkey: failed to register hotkey: CTRL + SHIFT + W")
			zap.L().Fatal(errBacklogBody.Error())
			return
		}
	}
	// TODO - Add a check for the environment variable to disable the hotkey
	// Register aged title hotkey
	errAgedTitle := pasteAgedTitle.Register()
	if errAgedTitle != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + E")
		zap.L().Fatal(errAgedTitle.Error())
		return
	}
	// Register aged body hotkey
	errAgedBody := pasteAgedBody.Register()
	if errAgedBody != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + SHIFT + E")
		zap.L().Fatal(errAgedBody.Error())
		return
	}
	// Register FDR title hotkey
	errFDRTitle := pasteFDRTitle.Register()
	if errFDRTitle != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + R")
		zap.L().Fatal(errFDRTitle.Error())
		return
	}
	// Register FDR body hotkey
	errFDRBody := pasteFDRBody.Register()
	if errFDRBody != nil {
		zap.L().Error("hotkey: failed to register hotkey: CTRL + SHIFT + R")
		zap.L().Fatal(errFDRBody.Error())
		return
	}

	// Run this on a infinite loop to watch for key events
	// Otherwise this doesn't watch further key events aside from the first one
	for {
		select {
		case <-pasteBacklogTitle.Keydown():
			zap.L().Info("hotkey: CTRL + W is down")
		// Create a new backlog title and paste it to the clipboard
		case <-pasteBacklogTitle.Keyup():
			backlogTitle := os.Getenv("BACKLOG_TITLE_PREFIX") + " | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(backlogTitle).Bytes())
			zap.L().Info("hotkey: CTRL + W is up")
		case <-pasteBacklogBody.Keydown():
			zap.L().Info("hotkey: CTRL + SHIFT + W is down")
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
			zap.L().Info("hotkey: CTRL + SHIFT + W is up")
		case <-pasteAgedTitle.Keydown():
			zap.L().Info("hotkey: CTRL + E is down")
		// Create a new backlog title and paste it to the clipboard
		case <-pasteAgedTitle.Keyup():
			agedTitle := os.Getenv("AGED_TITLE_PREFIX") + " | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(agedTitle).Bytes())
			zap.L().Info("hotkey: CTRL + E is up")
		case <-pasteAgedBody.Keydown():
			zap.L().Info("hotkey: CTRL + SHIFT + E is down")
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
			zap.L().Info("hotkey: CTRL + SHIFT + E is up")
		case <-pasteFDRTitle.Keydown():
			zap.L().Info("hotkey: CTRL + F is down")
		// Create a new backlog title and paste it to the clipboard
		case <-pasteFDRTitle.Keyup():
			FDRTitle := os.Getenv("FDR_TITLE_PREFIX") + " | " + time.Now().Format("2006/01/02")
			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(FDRTitle).Bytes())
			zap.L().Info("hotkey: CTRL + R is up")
		case <-pasteFDRBody.Keydown():
			zap.L().Info("hotkey: CTRL + SHIFT + R is down")
		// Read the file and paste it to the clipboard for the FDR body
		case <-pasteFDRBody.Keyup():
			FDRTitle := os.Getenv("FDR_BODY_FILE_PATH")
			FDRBodyContent, err := os.ReadFile(FDRTitle)
			if err != nil {
				zap.L().Error("failed to read file for FDR_BODY_FILE_PATH")
				zap.L().Error("is the filesytem read-only or does the file exist?")
				zap.L().Fatal(err.Error())
				return
			}

			clipboard.Write(clipboard.FmtText, bytes.NewBufferString(string(FDRBodyContent)).Bytes())
			zap.L().Info("hotkey: CTRL + SHIFT + R is up")
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
	app.Get("/", controllers.IndexController)
	app.Get("/api/v1/version", controllers.VersionController)
	// Run these functions in a goroutine
	// These keybind watchers are being watched on a infinite for loop. It seems better to run them in a goroutine because of this
	go func() {
		invokeHotKeys()
	}()
	// Change this to a not commonly used port to avoid issues with other local services
	fiberErr := app.Listen(":3090")

	if fiberErr != nil {
		zap.L().Fatal(fiberErr.Error())
	}
}
