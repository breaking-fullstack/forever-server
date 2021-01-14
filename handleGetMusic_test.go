package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/breaking-fullstack/forever-server/entity"
	"github.com/stretchr/testify/assert"
)

type getMusicTestService struct{}

func (*getMusicTestService) GetAll(context.Context, string) ([]entity.Music, error) {
	return getMusicTestList, nil
}

func TestHandleGetMusic(t *testing.T) {
	srv := NewServer("", &getMusicTestService{}, nil)
	testReq := httptest.NewRequest(http.MethodGet, "/music", nil)
	testRec := httptest.NewRecorder()

	srv.Handler.ServeHTTP(testRec, testReq)
	assert.Equal(t, 200, testRec.Code)

	expected, err := json.Marshal(getMusicTestList)
	assert.NoError(t, err)

	got, err := ioutil.ReadAll(testRec.Body)
	assert.NoError(t, err)

	assert.Equal(t, bytes.TrimSpace(expected), bytes.TrimSpace(got))
}

var getMusicTestList = []entity.Music{
	{
		Title: "Foo Bars",
		URL:   "http://foo.bars/music",
	},
	{
		Title: "Baz Bass",
		URL:   "http://baz.music",
	},
}

func (*getMusicTestService) Save(ctx context.Context, userID string, m entity.Music) error {
	return nil
}

func (*getMusicTestService) Delete(ctx context.Context, userID string, mID string) error { return nil }
