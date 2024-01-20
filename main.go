package main

import (
	"log"

	"github.com/baihaqism/bot-shopee/bot"
)

func main() {
	bot, err := bot.NewBot()
	if err != nil {
		log.Fatal(err)
	}

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
