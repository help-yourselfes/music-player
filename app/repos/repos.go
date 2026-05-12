package repos

import (
	"context"
	"music-player/app/models"
)

type TrackRepo interface {
	Create(ctx context.Context, track *models.Track) (int64, error)
	Read(ctx context.Context, id int64) (*models.Track, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*models.Track, error)
}

type PlaylistRepo interface {
	Create(ctx context.Context, playlist *models.Playlist) (int64, error)
	Read(ctx context.Context, id int64) (*models.Playlist, error)
	Delete(ctx context.Context, id int64) error
	ListTracks(ctx context.Context, id int64) ([]*models.Track, error)
	List(ctx context.Context) ([]*models.Playlist, error)
}
