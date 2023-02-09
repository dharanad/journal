package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Error setting up zap logger", err)
	}
	sugar := logger.Sugar()
	sugar.Info("Starting journal")

	r := gin.Default()
	r.GET("/ping", PingHandler())
	if err = r.Run(":80"); err != nil {
		sugar.Fatal("Error running http-server", err)
	}
}

func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
