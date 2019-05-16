package message

import (
	"github.com/girlvr/yinhe_bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/godcong/go-trait"
	"strings"
)

// Video ...
func Video(message *tgbotapi.Message) (ct []tgbotapi.Chattable) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	v := strings.Split(message.Text, WhiteSpace)

	if len(v) > 1 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索："+v[1]))
		if _, err := bot.Send(msg); err != nil {
			log.Error(err)
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
		ct = append(ct)
		return
	}

	if hasVideo {
		if _, err := bot.Send(config); err != nil {
			log.Error(err)
		}
		return
	}

	if _, err := bot.Send(msg); err != nil {
		return
	}
}
