package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRoot(t *testing.T) {
	expected := "https://github.com/breaking-fullstack/forever"

	testReq := httptest.NewRequest(http.MethodGet, "/", nil)
	testRec := httptest.NewRecorder()

	srv := NewServer(":8080", nil)
	srv.getRoutes().ServeHTTP(testRec, testReq)

	assert.Equal(t, http.StatusOK, testRec.Code)
	got, err := ioutil.ReadAll(testRec.Body)
	assert.NoError(t, err)

	gotStr := string(got)
	assert.Equal(t, expected, gotStr, "bad response from root handler")
}
