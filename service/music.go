package service

import (
	"context"

	"github.com/breaking-fullstack/forever-server/entity"
)

type MusicService interface {
	GetAll(ctx context.Context, userID string) ([]entity.Music, error)
	Save(ctx context.Context, userID string, m entity.Music) error
	Delete(ctx context.Context, userID string, mID string) error
}
