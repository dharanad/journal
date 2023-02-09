package handler

import "github.com/gin-gonic/gin"

func NewPingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
