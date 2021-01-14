package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/breaking-fullstack/forever-server/entity"
	"github.com/breaking-fullstack/forever-server/testhelper"
	"github.com/stretchr/testify/assert"
)

type getMusicTestService struct{}

func (*getMusicTestService) GetAll(context.Context, string) ([]entity.Music, error) {
	return testhelper.MusicList, nil
}

func TestHandleGetMusic(t *testing.T) {
	srv := NewServer("", &getMusicTestService{}, &testhelper.AuthVerifier{})
	testReq := httptest.NewRequest(http.MethodGet, "/music", nil)
	testReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", testhelper.ValidAuthJWT))
	testRec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(testRec, testReq)
	assert.Equal(t, 200, testRec.Code)

	expected, err := json.Marshal(testhelper.MusicList)
	assert.NoError(t, err)

	got, err := ioutil.ReadAll(testRec.Body)
	assert.NoError(t, err)

	assert.Equal(t, bytes.TrimSpace(expected), bytes.TrimSpace(got))
}

func (*getMusicTestService) Save(ctx context.Context, userID string, m entity.Music) error {
	return nil
}

func (*getMusicTestService) Delete(ctx context.Context, userID string, mID string) error { return nil }
