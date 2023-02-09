package handler

import (
	"github.com/dharanad/journal/internal/base"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	logger, _ := base.NewZapLogger()
	jr := base.NewJournalRouter(logger)
	router := jr.SetupRouter()
	jr.RegisterRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
