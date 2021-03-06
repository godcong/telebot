package api

import (
	"github.com/gin-gonic/gin"

	"github.com/godcong/telebot/database/ent"
)

func (a *api) handleMessage(context *gin.Context) {
	messages, err := a.db.QueryMessages(context.Request.Context())
	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(200, messages)
}

func (a *api) handleMessagePOST(context *gin.Context) {
	msg := new(ent.Message)
	var err error
	if err = context.BindJSON(msg); err != nil {
		context.JSON(500, err)
		return
	}
	if err = a.GetIntID(context, &msg.ID); err != nil {
		context.JSON(500, err)
		return
	}
	if msg, err = a.db.UpdateMessage(context.Request.Context(), msg); err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(200, msg)
}
