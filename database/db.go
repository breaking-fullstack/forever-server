package database

import (
	"context"

	"github.com/breaking-fullstack/forever-server/entity"
)

type DB interface {
	GetMusicList(ctx context.Context, userID string) ([]entity.Music, error)
	AddMusic(ctx context.Context, userID string, music entity.Music) error
	RemoveMusic(ctx context.Context, userID string, musicID string) error
}
