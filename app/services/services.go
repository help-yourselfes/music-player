package services

import (
	"context"
	"errors"
	"music-player/app/models"
	"music-player/app/repos"
)

type AppService struct {
	tracks    repos.TrackRepo
	playlists repos.PlaylistRepo
}

func NewAppService(tracks repos.TrackRepo, playlists repos.PlaylistRepo) *AppService {
	return &AppService{
		tracks:    tracks,
		playlists: playlists,
	}
}

func (s *AppService) AddTrack(ctx context.Context, reqTrack *models.Track) (int64, error) {
	if reqTrack.Path == "" || reqTrack.Name == "" {
		return -1, errors.New("no name")
	}

	id, err := s.tracks.Create(ctx, reqTrack)
	if err != nil {
		return -1, err
	}

	return id, nil
}
