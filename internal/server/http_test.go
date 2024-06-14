package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/data-exchange-csv/internal/server"
	"github.com/stretchr/testify/assert"
)

const port = "8080" // TODO: Replace with env variable
const endpoint = ":" + port

func TestHealthCheck(t *testing.T) {
	server.New() // TODO: Finish this

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/health", nil)

	assert.Equal(t, 200, w.Code)
}

func TestCSVValidation(t *testing.T) {

}
