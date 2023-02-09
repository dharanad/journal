package handler

import (
	"github.com/dharanad/journal/internal/base/handler"
	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
}

func NewPingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.HandleResponse(ctx, nil, PingResponse{Message: "Pong"})
	}
}
