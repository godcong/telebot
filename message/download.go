package message

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Download ...
func Download(message *tgbotapi.Message) (able tgbotapi.Chattable) {
	fb, e := getLocalFile("", GetProperty().Download)
	if e != nil {
		return tgbotapi.NewMessage(message.Chat.ID, "暂无更新文件!")
	}
	return tgbotapi.NewDocumentUpload(message.Chat.ID, fb)
}
