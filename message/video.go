package message

import (
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"strings"
)

// Video ...
func Video(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	photo := tgbotapi.NewPhotoUpload(message.Chat.ID, "")
	v := strings.Split(message.Text, " ")

	if len(v) > 1 {
		msg.Text = "正在搜索：" + v[1]
		if _, err := bot.Send(msg); err != nil {
			logrus.Error(err)
			msg.Text = "没有找到对应资源"
			return
		}
		video := searchVideo(v[1])
		if video == nil {
			msg.Text = "没有找到对应资源"
			return
		}

		e := parseVideo(&photo, video)
		if e != nil {
			msg.Text = "没有找到对应资源"
			return
		}
		_ = model.Visited(video)
		return
	}
}
