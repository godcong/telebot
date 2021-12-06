package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/godcong/telebot/abstract"
	"github.com/godcong/telebot/config"
	"github.com/godcong/telebot/database"
)

type api struct {
	*gin.Engine
	config *config.Config
	db     *database.DB
}

func (a *api) Run() error {
	g := a.Engine.Group("api/v0", a.middleware)
	g.Handle("GET", "statistics", a.handleStatistic)
	g.Handle("GET", "messages", a.handleMessage)
	g.Handle("POST", "message/:id", a.handleMessagePOST)

	go a.Engine.Run(":18080")

	return nil
}

func (a *api) GetIntID(context *gin.Context, i *int) error {
	id := context.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	*i = intID
	return nil
}

func New(cfg *config.Config, db *database.DB) abstract.API {
	engine := gin.Default()

	return &api{
		Engine: engine,
		config: cfg,
		db:     db,
	}
}

var _ abstract.API = (*api)(nil)
