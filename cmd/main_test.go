package main

import (
	"github.com/CDCgov/data-exchange-csv/cmd/internal/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDuplicateServer confirms if only one HTTP server can be active at any time
func TestDuplicateServer(t *testing.T) {
	err := server.New()
	assert.NotNil(t, err)
}
