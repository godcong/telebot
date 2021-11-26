package main

import (
	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/client"
)

func main() {
	c := client.NewClient(&config.Config{
		APIID:   "",
		APIHash: "",
	})
	c.Run()
}
