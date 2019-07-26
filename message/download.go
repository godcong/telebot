package message

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Download ...
func Download(message *tgbotapi.Message) (able tgbotapi.Chattable) {
	log.With("id", downloadFileID)
	if downloadFileID == "" {
		down := tgbotapi.NewDocumentUpload(message.Chat.ID, GetProperty().Download)
		downloadFileID = down.FileID
		return down
	}
	able = tgbotapi.NewDocumentUpload(message.Chat.ID, downloadFileID)
	return
}
