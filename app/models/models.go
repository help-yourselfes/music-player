package models

type PlaylistID = int64
type TrackID = int64

type Playlist struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Path          string    `json:"path"`
	PlaylistNames *[]string `json:"playlists"`
	Tracks        *[]Track
}

type Track struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
