package message

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// Download ...
func Download(message *tgbotapi.Message) (able tgbotapi.Chattable) {
	log.With("id", downloadFileID).Info("download")
	if downloadFileID == "" {
		return tgbotapi.NewDocumentUpload(message.Chat.ID, GetProperty().Download)
	}
	return tgbotapi.NewDocumentUpload(message.Chat.ID, downloadFileID)
}
