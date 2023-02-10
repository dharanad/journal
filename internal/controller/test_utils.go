package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func NewTestServer() *gin.Engine {
	r := gin.New()
	return r
}

func UnmarshalStruct(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
