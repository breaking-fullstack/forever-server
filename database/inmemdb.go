package database

import (
	"context"
	"errors"
	"sync"

	"github.com/breaking-fullstack/forever-server/entity"
)

//InMemDB is an in-memory implementation of the DB interface
type InMemDB struct {
	mu   sync.RWMutex
	data map[string][]entity.Music
}

func NewInMem() DB {
	return &InMemDB{
		data: make(map[string][]entity.Music),
	}
}

func (im *InMemDB) GetMusicList(ctx context.Context, userID string) ([]entity.Music, error) {
	im.mu.RLock()
	defer im.mu.RUnlock()
	return im.data[userID], nil
}

func (im *InMemDB) AddMusic(ctx context.Context, userID string, music entity.Music) error {
	im.mu.Lock()
	defer im.mu.Unlock()
	im.data[userID] = append(im.data[userID], music)
	return nil
}

func (im *InMemDB) RemoveMusic(ctx context.Context, userID string, musicID string) error {
	im.mu.Lock()
	defer im.mu.Unlock()
	userData := im.data[userID]
	var pos = -1
	for i, music := range userData {
		if music.ID == musicID {
			pos = i
			break
		}
	}
	if pos == -1 {
		return errors.New("music not found")
	}
	im.data[userID] = append(userData[:pos], userData[pos+1:]...)
	return nil
}
