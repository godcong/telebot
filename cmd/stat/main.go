package main

import (
	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/client"
)

func main() {
	c := client.NewClient(&config.Config{
		Bot:    config.Bot{},
		Client: config.Client{
			APIID:   "",
			APIHash: "",
		},
		Debug:  false,
		Auth:   "",
	})
	c.Run()
}
