package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/breaking-fullstack/forever-server/database"
	"github.com/breaking-fullstack/forever-server/service"
	"github.com/breaking-fullstack/forever-server/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestHandleAddMusic(t *testing.T) {
	srv := NewServer(":8080", service.NewMusic(database.NewInMem()), &testhelper.AuthVerifier{})

	testData := `
{
	"title": "Foo's Bars",
	"url": "http://foo.bar"
}
`
	testReq := httptest.NewRequest(http.MethodPost, "/music", strings.NewReader(testData))
	testReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testhelper.ValidAuthJWT))
	testRec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(testRec, testReq)
	assert.Equal(t, 201, testRec.Code)

	mlist, err := srv.musicService.GetAll(context.Background(), testhelper.TestUser)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(mlist))
}
