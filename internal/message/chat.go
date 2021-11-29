package message

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/database/ent"
	"github.com/motomototv/telebot/database/ent/message"
)

func getName(user *tgbotapi.User) string {
	if user.UserName != "" {
		return user.UserName
	}
	return user.LastName + "·" + user.FirstName
}

func actionChatMember(bot abstract.Bot, msg *ent.Message, update tgbotapi.Update) error {
	switch msg.Action {
	case message.ActionWelcome:
		var usrs []string
		for _, u := range *update.Message.NewChatMembers {
			usrs = append(usrs, getName(&u))
		}
		nmsg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(msg.Message, usrs))
		botMsg, err := bot.Bot().Send(nmsg)
		if err != nil {
			return err
		}
		if msg.AutoRemove {
			//t := time.AfterFunc(time.Duration(msg.AutoRemoveTime)*time.Second, func() {
			//	if _, err := bot.Bot().DeleteMessage(tgbotapi.DeleteMessageConfig{
			//		ChatID:    update.Message.Chat.ID,
			//		MessageID: botMsg.MessageID,
			//	}); err != nil {
			//		log.Println(err)
			//	}
			//})
			go func() {
				time.Sleep(time.Duration(msg.AutoRemoveTime) * time.Second)
				if _, err := bot.Bot().DeleteMessage(tgbotapi.DeleteMessageConfig{
					ChatID:    update.Message.Chat.ID,
					MessageID: botMsg.MessageID,
				}); err != nil {
					log.Println(err)
				}
			}()
		}
	case message.ActionStatistic:
		for _, u := range *update.Message.NewChatMembers {
			err := bot.DB().UpdateNewMemberStatistic(bot.Context(), &ent.Statistic{
				FirstName: u.FirstName,
				LatName:   u.LastName,
				FromUser:  update.Message.From.ID,
				ChannelID: update.Message.Chat.ID,
				UserID:    u.ID,
				JoinTime:  time.Now().Unix(),
			})
			if err != nil {
				fmt.Println("ERROR", "update statistic:", err)
				fmt.Sprintf("user:%+v\n", u)
			}
		}
	}
	return nil
}

func chatMemberWelcome() {

}
