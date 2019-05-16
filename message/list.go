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
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1"),
			tgbotapi.NewKeyboardButton("2"),
			tgbotapi.NewKeyboardButton("3"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("4"),
			tgbotapi.NewKeyboardButton("5"),
			tgbotapi.NewKeyboardButton("6"),
		),
	)

	if len(v) > 2 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索："+v[1]))
		//limit, _ := strconv.ParseUint(v[0], 10, 32)
		start, _ := strconv.ParseUint(v[0], 10, 32)
		videos, err := searchVideoList(10, int(start))
		if err != nil || videos == nil || len(videos) == 0 {
			ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
			return
		}
		addLine("资源列表" + v[0] + ":")
		for i, v := range videos {
			u := strconv.FormatUint(uint64(i), 10)
			msg.Text = u + ":" + v.Bangumi + " " + v.Intro
			if len(v.Role) > 0 {
				msg.Text += " " + v.Role[0]
			}
			msg.Text = addLine(msg.Text)
		}
		buttonMsg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		buttonMsg.ReplyMarkup = numericKeyboard
		ct = append(ct, msg)
	}
	return
}
