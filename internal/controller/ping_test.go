package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	r := NewTestServer()
	w := httptest.NewRecorder()
	r.GET("/ping", NewPingController().Ping())
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	actual := PingResponse{}
	err := UnmarshalStruct(w.Body.Bytes(), &actual)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, actual.Status)
	assert.Equal(t, "pong", actual.Message)
}
