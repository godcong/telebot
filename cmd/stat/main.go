package main

import (
	"flag"
	"fmt"

	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/client"
	"github.com/motomototv/telebot/log"
)

var path = flag.String("path", "bot.cfg", "default property path")

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}
	log.Debug = cfg.Debug
	c, err := client.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	c.Run()
	fmt.Println("end")
}
