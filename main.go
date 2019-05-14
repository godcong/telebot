package main

import (
	"fmt"
	"github.com/girlvr/yinhe_bot/message"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	token, e := ioutil.ReadFile("token")
	if e != nil {
		return
	}
	logrus.Info(string(token))
	logrus.SetReportCaller(true)
	BootWithGAE(string(token))
}

// BootWithGAE ...
func BootWithGAE(token string) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "443"
		log.Printf("Defaulting to port %s", port)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	t := "crVuYHQbUWCerib3"
	_, err = bot.SetWebhook(api.NewWebhook("https://bot.dhash.app/" + t))
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

	updates := bot.ListenForWebhook("/" + t)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("ping call")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("PONG"))
	})
	go http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "server.crt", "server.key", nil)
	message.InitBoot(bot)
	for update := range updates {
		message.HookMessage(update)
	}
}

// BootWithUpdate ...
func BootWithUpdate(token string) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		message.HookMessage(update)
	}
}
