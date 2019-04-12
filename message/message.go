package message

import (
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

var bot *tgbotapi.BotAPI

func InitBoot(botapi *tgbotapi.BotAPI) {
	bot = botapi
}

func HookMessage(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if !update.Message.IsCommand() { // ignore any non-command Messages
		return
	}

	// Create a new MessageConfig. We don't have text yet,
	// so we should leave it empty.
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	pto := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "")

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
			pto.File = "https://ipfs.io/ipfs/" + video.Poster
		}
	case "suggest":
		msg.Text = "result the top"
	case "help":
		msg.Text = "type /av or /suggest."
	case "status":
		msg.Text = "I'm ok."

	}
	if _, err := bot.Send(pto); err != nil {
		log.Panic(err)
	}
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func searchVideo(s string) *model.Video {
	video := &model.Video{}
	if b, err := model.FindVideo(s, video); err != nil || !b {
		return nil
	}
	return video
}
