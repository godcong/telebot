package message

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// List ...
func List(message *tgbotapi.Message) (ct []tgbotapi.Chattable) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	v := strings.Split(message.Text, WhiteSpace)
	ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "正在搜索..."))

	numStr := "0"
	start := uint64(0)
	if len(v) > 1 {
		//limit, _ := strconv.ParseUint(v[0], 10, 32)
		start, _ = strconv.ParseUint(v[1], 10, 32)
		numStr = v[1]
	}
	limit := uint64(10)
	if len(v) > 2 {
		limit, _ = strconv.ParseUint(v[2], 10, 32)
		if limit > 25 {
			limit = 25
		}
	}

	videos, err := searchVideoList(int(limit), int(start))
	if err != nil || videos == nil || len(videos) == 0 {
		ct = append(ct, tgbotapi.NewMessage(message.Chat.ID, "没有找到对应资源"))
		return
	}
	next := fmt.Sprintf("/list %d", start+limit)
	if len(videos) < int(limit) {
		next = "/close"
	}
	pre := "/close"
	if start >= limit {
		pre = fmt.Sprintf("/list %d", start-limit)
	} else if start > 0 && start < limit {
		pre = fmt.Sprintf("/list %d", 0)
	} else {
		// close
	}
	row := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(pre),
		tgbotapi.NewKeyboardButton(next),
	)
	numericKeyboard := tgbotapi.NewReplyKeyboard(row)
	msg.Text = addLine("资源列表" + numStr + ":")
	for i, v := range videos {
		msg.Text += fmt.Sprintf("%d. [%s] %s ☆Hot %d☆", i+1, v.Bangumi, v.Intro, v.Visit)
		msg.Text = addLine(msg.Text)
	}
	ct = append(ct, msg)
	buttonMsg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	buttonMsg.ReplyMarkup = numericKeyboard
	ct = append(ct, buttonMsg)
	return
}
