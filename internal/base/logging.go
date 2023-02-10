package base

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitLogging() {
	gin.SetMode(gin.ReleaseMode)
}

func NewZapLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func NewSugarLogger() (*zap.SugaredLogger, error) {
	logger, err := NewZapLogger()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
