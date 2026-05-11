package models

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
