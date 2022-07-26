package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandleHello(t *testing.T) {
	s := new(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
