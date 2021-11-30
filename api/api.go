package api

import (
	"github.com/gin-gonic/gin"

	"github.com/motomototv/telebot/abstract"
	"github.com/motomototv/telebot/config"
	"github.com/motomototv/telebot/database"
)

type api struct {
	*gin.Engine
	config *config.Config
	db     *database.DB
}

func (a *api) Run() error {
	g := a.Engine.Group("api/v0", a.middleware)
	g.Handle("GET", "statistics", a.handleStatistic)

	go a.Engine.Run(":18080")

	return nil
}

func (a *api) handleStatistic(context *gin.Context) {
	statistics, err := a.db.QueryStatistics(context.Request.Context())
	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(200, statistics)
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
