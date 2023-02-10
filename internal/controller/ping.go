package controller

import (
	"github.com/dharanad/journal/internal/base/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

type PingResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
}

func (c *PingController) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.HandleResponse(ctx, nil, PingResponse{Message: "pong", Status: http.StatusOK})
	}
}
