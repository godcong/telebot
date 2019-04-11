package main

import (
	"fmt"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strings"
)

func BootWithGAE(token string) {
	bot, err := api.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	t := "thisistoken"
	_, err = bot.SetWebhook(api.NewWebhook("https://yinhe-bot.appspot.com/" + t))
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
	go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := api.NewMessage(update.Message.Chat.ID, "")
		var photo api.PhotoConfig

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "av":
			v := strings.Split(update.Message.Text, " ")
			msg.Text = "av not found"
			if len(v) > 1 {
				video := searchVideo(v[1])
				if video == nil {
					return
				}
				msg.Text = ""
				for _, value := range video.VideoGroupList {
					for _, o := range value.Object {
						msg.Text += "https://ipfs.io/ipfs/" + o.Link.Hash + "\n"
					}
				}
				photo = api.NewPhotoUpload(update.Message.Chat.ID, "https://ipfs.io/ipfs/"+video.Poster)
			}
		case "suggest":
			msg.Text = "result the top"
		case "help":
			msg.Text = "type /av or /suggest."
		case "status":
			msg.Text = "I'm ok."

		}
		if _, err := bot.Send(photo); err != nil {
			log.Panic(err)
		}
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

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
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we should leave it empty.
		msg := api.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "type /sayhi or /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "av":
			v := strings.Split(update.Message.Text, " ")
			msg.Text = "av not found"
			if len(v) > 1 {
				msg.Text = searchVideo(v[1])
			}
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
