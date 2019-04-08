package main

import (
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

var token = flag.String("token", "", "set new bot api token")

func main() {
	flag.Parse()
	bot, err := tgbotapi.NewBotAPI(*token)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	webHook := tgbotapi.NewWebhookWithCert("https://y11e.com/"+bot.Token, "/etc/letsencrypt/live/y11e.com/fullchain.pem")
	//webHook := tgbotapi.NewWebhook("https://52.194.147.237/ping/" + bot.Token)
	_, err = bot.SetWebhook(webHook)
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	fmt.Println("telegram bot starting")
	//updates := bot.ListenForWebhook("/" + bot.Token)

	http.ListenAndServe(":8080", NewRouter(bot.Token))

	//for update := range updates {
	//	log.Printf("%+v\n", update)
	//}
	//http.ListenAndServe(":8080", NewRouter())

}
