package api

import (
	"github.com/gin-gonic/gin"
)

func (a *api) handleStatistic(context *gin.Context) {
	statistics, err := a.db.QueryStatistics(context.Request.Context())
	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(200, statistics)
}
