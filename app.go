package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Playlist struct {
	Path   string `json:"path"`
	Tracks *[]Track
}
type Track struct {
	Name string `json:"name"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetTracks(path string) *Playlist {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var tracks []Track

	for _, file := range files {
		tracks = append(tracks, Track{
			Name: file.Name(),
		})
	}

	return &Playlist{
		Path:   path,
		Tracks: &tracks,
	}
}
