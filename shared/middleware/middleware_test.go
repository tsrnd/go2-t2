package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/infrastructure"
)

func TestLoggerMiddleware(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Add("Accept-Language", "en_US")
	rr := httptest.NewRecorder()
	Logger(infrastructure.NewLogger())(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestHeaderMiddleware(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	Header(infrastructure.NewLogger())(handler).ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
