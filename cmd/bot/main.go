package main

import (
	"flag"

	_ "github.com/mattn/go-sqlite3"

	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/bot"
)

var path = flag.String("path", "yinhe.json", "default property path")
//var port = flag.String("port", "443", "default port")
//
func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}
	telebot, err := bot.NewBot(cfg)
	if err != nil {
		panic(err)
	}
	if err := telebot.Run(); err != nil {
		panic(err)
	}
}
