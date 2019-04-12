package message

import (
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const ServerURL = "https://ipfs.io/ipfs/"

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
	case "video":
		v := strings.Split(update.Message.Text, " ")
		msg.Text = "video not found"
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
			pto.File = getFile(video.Poster)
		}
	case "top":
		msg.Text = "result the top"
	case "help":
		msg.Text = "type /video or /top."
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

func getFile(hash string) string {
	resp, err := http.Get(url(hash))
	if err != nil {
		logrus.Error(err)
		return ""
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	err = ioutil.WriteFile(hash, bytes, os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return hash
}

func url(hash string) string {
	return ServerURL + hash
}
