package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorRespBody struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func NewRespBodyFromError(code int, reason string) *ErrorRespBody {
	return &ErrorRespBody{Code: code, Reason: reason}
}

func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, data)
		return
	}
	ctx.JSON(http.StatusInternalServerError, NewRespBodyFromError(http.StatusInternalServerError, err.Error()))
}
