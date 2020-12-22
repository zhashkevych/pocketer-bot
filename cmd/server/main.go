package main

import (
	"github.com/zhashkevych/telegram-pocket-bot/pkg/server"
	"log"
)

func main() {
	redirectServer := server.NewRedirectServer()

	if err := redirectServer.Start(); err != nil {
		log.Fatal(err)
	}
}
