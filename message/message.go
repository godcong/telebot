package message

import (
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const ServerURL = "https://ipfs.io/ipfs/"

var bot *tgbotapi.BotAPI
var photoHas = make(map[string][]byte)

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
	hasVideo := false
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
			pto.File = getFile(video.Poster)
			for _, value := range video.VideoGroupList {
				for _, o := range value.Object {
					hasVideo = true
					pto.Caption += url(o.Link.Hash) + "\n"
				}
			}
		}
	case "top":
		msg.Text = "result the top"
	case "help":
		msg.Text = "type /video or /top."
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "type /video or /top."
	}

	if hasVideo {
		if _, err := bot.Send(pto); err != nil {
			log.Panic(err)
		}
		return
	}
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func searchVideo(s string) *model.Video {
	video := &model.Video{}
	if b, err := model.DeepFind(s, video); err != nil || !b {
		return nil
	}
	return video
}

func getFile(hash string) tgbotapi.FileBytes {
	url := url(hash)
	logrus.Info("url:", url)
	fb := tgbotapi.FileBytes{
		Name: hash,
	}
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return fb
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return fb
	}
	fb.Bytes = bytes
	return fb
}

func url(hash string) string {
	return ServerURL + hash
}
