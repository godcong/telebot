package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/godcong/telebot/api"
	"github.com/godcong/telebot/config"
	"github.com/godcong/telebot/internal/bot"
	"github.com/godcong/telebot/log"
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
	log.Debug = cfg.Debug
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

	_, ok := <-interrupts
	if ok {
		fmt.Println("interrupt exit")
	}
	return nil
}
