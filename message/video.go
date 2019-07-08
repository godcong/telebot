package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yinhevr/seed/model"
	"strings"
)

// Video ...
func Video(message *tgbotapi.Message) (ct []tgbotapi.Chattable) {
	photo := tgbotapi.NewPhotoUpload(message.Chat.ID, "")
	v := strings.Split(message.Text, WhiteSpace)

	if len(v) > 1 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索："+v[1]))
		videos := searchVideo(strings.ToUpper(v[1]))
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
		ct = append(ct, photo)
	}
	return
}
