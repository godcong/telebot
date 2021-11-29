package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/database/ent/schema"
)

const help = `输入:
/v 或 /video +番号 查询视频 如：/v ABP-785
/t 或 /top　显示推荐视频
/l 或 /list 显示列表（仅私聊）点击:@yinhe_bot，私聊获取更多信息
/d 或 /down 获取求哈希下载链接（仅私聊）
/r 或 /rule 查看群规（仅私聊）
/h 或 /help 显示帮助`

func Command(bot abstract.Bot, msgT schema.MessageType, update tgbotapi.Update) error {
	//update.Message.Command()
	return nil
}
