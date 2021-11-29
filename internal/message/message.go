package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/database/ent"
	"github.com/motomototv/telebot/database/ent/schema"
	"github.com/motomototv/telebot/log"
)

var actions = [schema.MessageTypeMax]func(bot abstract.Bot, message *ent.Message, update tgbotapi.Update) error{
	schema.MessageTypeNone:       actionNone,
	schema.MessageTypeMessage:    actionMessage,
	schema.MessageTypeChatMember: actionChatMember,
}

func actionNone(bot abstract.Bot, message *ent.Message, update tgbotapi.Update) error {
	return nil
}

func actionMessage(bot abstract.Bot, message *ent.Message, update tgbotapi.Update) error {
	log.Println("received new message from:", update.Message.From.ID, "channel:", update.Message.Chat.ID)
	if update.Message.Chat.IsPrivate() {
		log.Println("skip statistic with private talk")
		return nil
	}
	return bot.DB().UpdateChatStatistic(bot.Context(), &ent.Statistic{
		FirstName: update.Message.From.FirstName,
		LatName:   update.Message.From.LastName,
		UserName:  update.Message.From.UserName,
		UserID:    update.Message.From.ID,
		FromUser:  0,
		ChannelID: update.Message.Chat.ID,
	})
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
		log.Println("Process new message:", schema.MessageType(message.Type))
		log.Printfln("Update message from detail:%+v", update.Message.From)
		log.Printfln("Update message chat detail:%+v", update.Message.Chat)
		if err := actions[message.Type](bot, message, update); err != nil {
			log.Println("ERROR:", "type:", message.Type, "error:", err)
		}
	}
	return nil
}
