package base

import (
	"github.com/dharanad/journal/internal/handler"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type JournalRouter struct {
	logger *zap.Logger
}

func NewJournalRouter(logger *zap.Logger) *JournalRouter {
	return &JournalRouter{logger: logger}
}

func (s *JournalRouter) SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(s.logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(s.logger, true))
	return r
}

func (s *JournalRouter) RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", handler.NewPingHandler())
}
