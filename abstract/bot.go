package abstract

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/godcong/telebot/config"
	"github.com/godcong/telebot/database"
)

type Bot interface {
	Run() error
	Bot() *tgbotapi.BotAPI
	DB() *database.DB
	Config() *config.Config
	Context() context.Context
}
