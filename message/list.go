package message

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

// List ...
func List(message *tgbotapi.Message) (ct []tgbotapi.Chattable) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	v := strings.Split(message.Text, WhiteSpace)

	ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索..."))
	start := uint64(0)
	row := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/list 10"),
	)
	numStr := "0"
	if len(v) > 1 {
		//limit, _ := strconv.ParseUint(v[0], 10, 32)
		start, _ = strconv.ParseUint(v[1], 10, 32)
		next := fmt.Sprintf("/list %d", start+10)
		if start >= 10 {
			pre := fmt.Sprintf("/list %d", start-10)
			row = tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(pre),
				tgbotapi.NewKeyboardButton(next),
			)
		}
		numStr = v[1]
	}

	numericKeyboard := tgbotapi.NewReplyKeyboard(row)

	videos, err := searchVideoList(10, int(start))
	if err != nil || videos == nil || len(videos) == 0 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
		return
	}
	msg.Text = addLine("资源列表" + numStr + ":")
	for i, v := range videos {
		u := strconv.FormatUint(uint64(i), 10)
		msg.Text += u + ". [" + v.Bangumi + "] " + v.Intro
		if len(v.Role) > 0 {
			msg.Text += " " + v.Role[0]
		}
		msg.Text = addLine(msg.Text)
	}
	ct = append(ct, msg)
	buttonMsg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	buttonMsg.ReplyMarkup = numericKeyboard
	ct = append(ct, buttonMsg)
	return
}
