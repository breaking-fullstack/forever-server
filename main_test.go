package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunAddr(t *testing.T) {
	testPort := "1025"
	os.Setenv("PORT", testPort)
	got := getRunAddr()
	assert.Equal(t, ":1025", got, "getRunAddr does not read port from environment")

	os.Unsetenv("PORT")
	got = getRunAddr()
	assert.Equal(t, ":8080", got, "getRunAddr does not default to port 8080")
}
