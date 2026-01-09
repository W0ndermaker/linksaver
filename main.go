package main

import (
	"flag"
	"linksaver/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// token = flags.Get(token)
	// token := mustToken()

	tgClient := telegram.New(tgBotHost, mustToken()) // телеграмм-клиент

	// fetcher = fetcher.New() --- получает новые события

	// processor = processor.New() --- после обработки отправит новые сообщения

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
}
