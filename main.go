package main

import (
	"flag"
	"log"

	tgClient "linksaver/clients/telegram"
	eventconsumer "linksaver/consumer/event-consumer"
	tg "linksaver/events/telegram2"
	"linksaver/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100 // Limit of events arriving at the same time
)

func main() {
	tgClient := tgClient.New(tgBotHost, mustToken()) // Telegram Client is responsible for HTTP requests(getting and sending events) to Telegram API

	eventsProcessor := tg.New(tgClient, files.New(storagePath)) // Fetch and Process events and also saves URLs in storage
	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize) // Consumer starts the program

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

// Parse args to get Telegram Bot token
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to TG bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
