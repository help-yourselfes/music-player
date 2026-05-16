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
func (s *AppService) AddPlaylist(ctx context.Context, name string, path string) (int64, error) {
	if name == "" {
		return -1, errors.New("no name")
	}

	playlist := &models.Playlist{
		Name: name,
		Path: path,
	}

	id, err := s.playlists.Create(ctx, playlist)
	if err != nil {
		return -1, err
	}

	playlist.ID = id

	return id, nil
}

func (s *AppService) GetPlaylistTracks(ctx context.Context, id models.PlaylistID) ([]*models.Track, error) {
	list, err := s.playlists.ListTracks(ctx, id)
	return list, err
}

func (s *AppService) AddPlaylistTrack(ctx context.Context, playlistID int64, trackID int64) error {
	return s.playlists.AddTrack(ctx, playlistID, trackID)
}

func (s *AppService) RemovePlaylistTrack(ctx context.Context, playlistID int64, trackID int64) error {
	return s.playlists.RemoveTrack(ctx, playlistID, trackID)
}
