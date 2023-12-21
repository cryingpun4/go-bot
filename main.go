package main

import (
	"flag"
	"log"
)

func main() {
	mustToken()
}

func mustToken() string {
	token := flag.String(
		"token-bot",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
