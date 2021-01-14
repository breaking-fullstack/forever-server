package service

import (
	"context"
	"testing"

	"github.com/breaking-fullstack/forever-server/database"
	"github.com/breaking-fullstack/forever-server/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestMusicService(t *testing.T) {
	testMusicServiceImplementation(t, &musicServiceImpl{database.NewInMem()})
}

func testMusicServiceImplementation(t *testing.T, ms Music) {
	testUID := "testUID"

	//Add test music
	for _, m := range testhelper.MusicList {
		err := ms.Save(context.Background(), testUID, m)
		assert.NoError(t, err)
	}

	//Delete last music inserted
	lastMusicID := testhelper.MusicList[len(testhelper.MusicList)-1].ID
	err := ms.Delete(context.Background(), testUID, lastMusicID)
	assert.NoError(t, err)

	//Verify changes
	savedList, err := ms.GetAll(context.Background(), testUID)
	assert.NoError(t, err)

	assert.Equal(t, len(testhelper.MusicList)-1, len(savedList))
	assert.Equal(t, testhelper.MusicList[0], savedList[0])
}
