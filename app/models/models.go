package models

type PlaylistID int64
type TrackID int64

type Playlist struct {
	ID            PlaylistID `json:"id"`
	Name          string     `json:"name"`
	Path          string     `json:"path"`
	PlaylistNames *[]string  `json:"playlists"`
	Tracks        *[]Track
}

type Track struct {
	ID   TrackID `json:"id"`
	Name string  `json:"name"`
	Path string  `json:"path"`
}
