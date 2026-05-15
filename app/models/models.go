package models

type playlistID int64
type trackID int64

type Playlist struct {
	ID            playlistID `json:"id"`
	Name          string     `json:"name"`
	Path          string     `json:"path"`
	PlaylistNames *[]string  `json:"playlists"`
	Tracks        *[]Track
}

type Track struct {
	ID   trackID `json:"id"`
	Name string  `json:"name"`
	Path string  `json:"path"`
}
