package server

import (
	"github.com/dharanad/journal/internal/controller"
	"github.com/dharanad/journal/internal/router"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.New()
	journalRouter := router.NewJournalRouter(controller.NewPingController())
	journalApiGroup := r.Group("/journal/api/v1")
	journalRouter.RegisterAPIRouter(journalApiGroup)
	return r
}
