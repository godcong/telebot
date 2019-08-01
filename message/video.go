package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yinhevr/seed/model"
	"strings"
)

// Video ...
func Video(message *tgbotapi.Message, s string) (ct []tgbotapi.Chattable) {
	photo := tgbotapi.NewPhotoUpload(message.Chat.ID, "")

	videos := searchVideo(strings.ToUpper(s))
	if videos == nil || len(videos) <= 0 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
		return
	}
	e := parseVideoInfo(&photo, videos)
	if e != nil {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
		return
	}
	_ = model.Visited(videos[0])
	return append(ct, photo)
}
