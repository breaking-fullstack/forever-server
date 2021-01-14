package service

import (
	"context"

	"github.com/breaking-fullstack/forever-server/database"
	"github.com/breaking-fullstack/forever-server/entity"
)

type Music interface {
	GetAll(ctx context.Context, userID string) ([]entity.Music, error)
	Save(ctx context.Context, userID string, m entity.Music) error
	Delete(ctx context.Context, userID string, mID string) error
}

type musicServiceImpl struct {
	db database.DB
}

func NewMusic(db database.DB) Music {
	return &musicServiceImpl{
		db: db,
	}
}

func (ms *musicServiceImpl) GetAll(ctx context.Context, userID string) ([]entity.Music, error) {
	return ms.db.GetMusicList(ctx, userID)
}

func (ms *musicServiceImpl) Save(ctx context.Context, userID string, m entity.Music) error {
	return ms.db.AddMusic(ctx, userID, m)
}

func (ms *musicServiceImpl) Delete(ctx context.Context, userID string, mID string) error {
	return ms.db.RemoveMusic(ctx, userID, mID)
}
