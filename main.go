package main

import (
	"bot/bot"
	"fmt"
	"os"
)

func main() {
	config, err := bot.ParseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config parse err: %s", err)
		os.Exit(1)
	}
	b, err := bot.NewBot(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bot init err: %s", err)
		os.Exit(1)
	}
	b.Run()
}