package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

// List ...
func List(message *tgbotapi.Message) (ct []tgbotapi.Chattable) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	v := strings.Split(message.Text, WhiteSpace)

	if len(v) > 2 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索："+v[1]))
		//limit, _ := strconv.ParseUint(v[0], 10, 32)
		start, _ := strconv.ParseUint(v[0], 10, 32)
		videos, err := searchVideoList(10, int(start))
		if err != nil || videos == nil {
			ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
			return
		}
		addLine("资源列表" + v[0] + ":")
		e := parseVideoBan(&msg, videos)
		if e != nil {
			ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
			return
		}

		ct = append(ct, msg)
	}
	return
}
