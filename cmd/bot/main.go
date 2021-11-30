package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/motomototv/telebot/api"
	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/bot"
)

var path = flag.String("path", "config", "default property path")

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
	api := api.New(cfg, telebot.DB())
	if err := api.Run(); err != nil {
		panic(err)
	}

	handleInterrupt()
}

func handleInterrupt() error {
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupts:
		fmt.Println("interrupt exit")
		return nil
	}

}
