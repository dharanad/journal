package main

import (
	"github.com/dharanad/journal/internal/base"
	"github.com/dharanad/journal/internal/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	logger, err := base.NewZapLogger()
	defer logger.Sync()
	if err != nil {
		log.Fatal("Error setting up zap logger", err)
	}
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.GET("/ping", handler.NewPingHandler())
	if err := r.Run(":80"); err != nil {
		panic("Error starting web server. Reason :" + err.Error())
	}
}
