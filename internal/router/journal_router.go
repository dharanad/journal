package router

import (
	"github.com/dharanad/journal/internal/controller"
	"github.com/gin-gonic/gin"
)

type JournalRouter struct {
	pingController *controller.PingController
}

func NewJournalRouter(pingController *controller.PingController) *JournalRouter {
	return &JournalRouter{pingController: pingController}
}

func (c *JournalRouter) RegisterAPIRouter(r *gin.RouterGroup) {
	r.GET("/ping", c.pingController.Ping())
}
