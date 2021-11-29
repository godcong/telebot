package api

import (
	"github.com/gin-gonic/gin"

	"github.com/motomototv/telebot/abstract"
)

type api struct {
	*gin.Engine
}

func (api) Run() error {
	return nil
}

func New() abstract.API {
	engine := gin.Default()
	return &api{
		Engine: engine,
	}
}

var _ abstract.API = (*api)(nil)
