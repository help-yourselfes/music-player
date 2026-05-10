package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Playlist struct {
	Name          string    `json:"name"`
	Path          string    `json:"path"`
	PlaylistNames *[]string `json:"playlists"`
	Tracks        *[]Track
}

type Track struct {
	Name string `json:"name"`
}

// App struct
type App struct {
	formats    []string
	formatsSet map[string]struct{}
	ctx        context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		formats: []string{"mp3", "opus", "ogg", "flac"},
	}
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

func (a *App) getFormats() map[string]struct{} {
	if len(a.formatsSet) == 0 {
		a.formatsSet = make(map[string]struct{})

		for _, ext := range a.formats {
			a.formatsSet[ext] = struct{}{}
		}

	}

	return a.formatsSet
}

func (a *App) GetTracks(path string) *Playlist {
	rootName := filepath.Base(path)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var playlists []string
	var tracks []Track

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		name := file.Name()
		// ext := filepath.Ext(name)

		if info.IsDir() {
			playlists = append(playlists, rootName)
			continue
		}

		if name[0] == '.' {
			continue
		}

		tracks = append(tracks, Track{
			Name: name,
		})
	}

	return &Playlist{
		Name:          rootName,
		Path:          path,
		PlaylistNames: &playlists,
		Tracks:        &tracks,
	}
}
