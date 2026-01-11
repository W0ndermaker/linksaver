package main

import (
	"flag"
	tgClient "linksaver/clients/telegram"
	eventconsumer "linksaver/consumer/event-consumer"
	tg "linksaver/events/telegram2"
	"linksaver/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	// token = flags.Get(token)

	// token := mustToken()

	// fetcher = fetcher.New() --- получает новые события из API TG

	// processor = processor.New() --- после обработки отправит новые сообщения в TG

	tgClient := tgClient.New(tgBotHost, mustToken()) // телеграмм-клиент

	eventsProcessor := tg.New(tgClient, files.New(storagePath))

	log.Print("service started")

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to TG bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
