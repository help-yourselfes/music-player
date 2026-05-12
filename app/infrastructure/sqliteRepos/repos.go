package sqliteRepos

import (
	"context"
	"database/sql"
	"music-player/app/models"
	"music-player/app/repos"
)

type TrackRepoSQLite struct {
	db *sql.DB
}

var _ repos.TrackRepo = &TrackRepoSQLite{}

type tr = TrackRepoSQLite

func NewTrackRepoSQLite(db *sql.DB) *TrackRepoSQLite {
	return &TrackRepoSQLite{db: db}
}

func (r *tr) Create(ctx context.Context, track *models.Track) (int64, error) {
	query := `
	INSERT OR IGNORE INTO tracks (
		name, 
		path
	) VALUES (?, ?)	`
	res, err := r.db.ExecContext(ctx, query, track.Name, track.Path)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, err
}

func (r *tr) Read(ctx context.Context, id int64) (*models.Track, error) {
	query := `
	SELECT * FROM tracks
	WHERE id = ?`
	res := r.db.QueryRowContext(ctx, query, id)

	var track models.Track

	err := res.Scan(
		&track.ID,
		&track.Name,
		&track.Path,
	)

	if err != nil {
		// custom errors check
		return nil, err
	}

	return &track, nil
}

func (r *tr) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM tracks
	WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *tr) List(ctx context.Context) ([]*models.Track, error) {
	query := `
	SELECT * FROM tracks`
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tracks := make([]*models.Track, 0)
	for rows.Next() {
		var track models.Track

		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Path,
		)

		if err != nil {
			return nil, err
		}

		tracks = append(tracks, &track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, err
}

type PlaylistRepoSQLite struct {
	db *sql.DB
}

func NewPlaylistRepoSQLite(db *sql.DB) *PlaylistRepoSQLite {
	return &PlaylistRepoSQLite{db: db}
}

var _ repos.PlaylistRepo = &PlaylistRepoSQLite{}

type pl = PlaylistRepoSQLite

func (r *pl) Create(ctx context.Context, playlist *models.Playlist) (int64, error) {
	query := `
	INSERT OR IGNORE INTO playlists (
		name
	) VALUES (?)`
	res, err := r.db.ExecContext(ctx, query, playlist.Name)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, err
}

func (r *pl) Read(ctx context.Context, id int64) (*models.Playlist, error) {
	query := `
	SELECT * FROM playlists
	WHERE id = ?`
	res := r.db.QueryRowContext(ctx, query, id)

	var playlist models.Playlist

	err := res.Scan(
		&playlist.ID,
		&playlist.Name,
	)

	if err != nil {
		// custom errors check
		return nil, err
	}

	return &playlist, nil
}

func (r *pl) Delete(ctx context.Context, id int64) error {
	query := `
	DELETE FROM playlists
	WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *pl) ListTracks(ctx context.Context, id int64) ([]*models.Track, error) {
	query := `
	SELECT * FROM tracks
	JOIN playlist_tracks ON tracks.id = playlist_tracks.trackID
	WHERE playlist_tracks.playlistID = ?`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tracks := make([]*models.Track, 0)
	for rows.Next() {
		var track models.Track

		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.Path,
		)

		if err != nil {
			return nil, err
		}

		tracks = append(tracks, &track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, err
}

func (r *pl) List(ctx context.Context) ([]*models.Playlist, error) {
	query := `
	SELECT * FROM playlists`
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playlists := make([]*models.Playlist, 0)
	for rows.Next() {
		var playlist models.Playlist

		err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
		)

		if err != nil {
			return nil, err
		}

		playlists = append(playlists, &playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return playlists, err
}
