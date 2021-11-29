package message

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/database/ent"
	"github.com/motomototv/telebot/database/ent/schema"
)

var actions = [schema.MessageTypeMax]func(bot abstract.Bot, message *ent.Message, update tgbotapi.Update) error{
	schema.MessageTypeChatMember: chatMember,
}

func Message(bot abstract.Bot, msgT schema.MessageType, update tgbotapi.Update) error {
	messages, err := bot.DB().QueryMessages(bot.Context(), msgT)
	if err != nil {
		return err
	}
	for _, message := range messages {
		if schema.MessageType(message.Type) >= schema.MessageTypeMax {
			continue
		}
		if err := actions[message.Type](bot, message, update); err != nil {
			fmt.Println("ERROR:", "type:", message.Type, "error:", err)
		}
	}
	return nil
}
