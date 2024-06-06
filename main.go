package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	godotenv "github.com/joho/godotenv"
	"github.com/zarldev/go-base/ui"
)

func main() {
	loadEnv()
	fmt.Println("Welcome to the program.")
	// start the program
	cfg := getConfig()
	go ui.StartUI(context.Background(), cfg)
	// create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)
	// register the channel to receive SIGINT signals
	signal.Notify(sigs, syscall.SIGINT)
	fmt.Println("Press Ctrl+C to stop the program")
	// wait for a SIGINT signal (Ctrl+C)
	<-sigs
	fmt.Println("Program stopped.")
}

func getConfig() ui.Config {
	return ui.Config{
		Port: os.Getenv("PORT"),
	}
}

func loadEnv() {
	// load environment variables
	// you can
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
