package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/internal/client"
	"github.com/motomototv/telebot/log"
)

var path = flag.String("path", "bot.cfg", "default property path")
var chatid = flag.Int64("chatid", 0, "chat id")

func main() {
	flag.Parse()
	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}
	log.Debug = cfg.Debug
	c, err := client.NewClient("user1", cfg)
	if err != nil {
		panic(err)
	}
	go c.Run()
	fmt.Println("Bot is running")

	members, err := c.SearchChatMembersByID(*chatid)
	if err == nil {
		fmt.Println("Group members:", len(members.Members), members.TotalCount)
		for i := range members.Members {

			fmt.Println("User:", members.Members[i].MemberID)
			request, err := c.GetUserByID(members.Members[i].MemberID)
			if err == nil {
				fmt.Println("User:", request.Username, "joined chat:", -1102281440)
			}

		}
	} else {
		fmt.Println("GetChatMembers error:", err)
	}

	handleInterrupt()
	fmt.Println("end")
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
