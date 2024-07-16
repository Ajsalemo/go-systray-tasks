package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	// Notify the application of the below signals to be handled on shutdown
	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	// Goroutine to clean up prior to shutting down
	go func() {
		sig := <-s
		switch sig {
		case os.Interrupt:
			fmt.Println("Caught SIGINT")
			systray.Quit()
		case syscall.SIGTERM:
			fmt.Println("Caught SIGTERM")
			systray.Quit()
		case syscall.SIGQUIT:
			fmt.Println("Caught SIGQUIT")
			systray.Quit()
		case syscall.SIGINT:
			fmt.Println("Caught SIGINT")
			systray.Quit()
		}
	}()

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/task.ico"))
	systray.SetTitle("I'm alive!")
	systray.SetTooltip("Look at me, I'm a tooltip!")
}

func onExit() {
	// Cleaning stuff here.
	systray.Quit()
	time.Sleep(2 * time.Second)
	fmt.Println("Exiting...")
	os.Exit(0)
}

func getIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
