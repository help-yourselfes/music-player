package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"music-player/app/infrastructure/filecheck"
	"music-player/app/models"
	"music-player/app/services"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	formats    []string
	formatsSet map[string]struct{}

	service services.AppService
	ctx     context.Context
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

	err := a.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to storage:\t", err)
	}
}

func (a *App) getFormats() map[string]struct{} {
	if len(a.formatsSet) == 0 {
		a.formatsSet = make(map[string]struct{})

		for _, ext := range a.formats {
			a.formatsSet["."+ext] = struct{}{}
		}

	}

	return a.formatsSet
}

func (a *App) GetTracks(path string) *models.Playlist {
	rootName := filepath.Base(path)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var playlists []string
	var tracks []models.Track

	for _, file := range files {
		if !filecheck.IsEntryTrack(file) {
			continue
		}

		name := file.Name()
		tracks = append(tracks, models.Track{
			Name: name,
			Path: path + "/" + name,
		})
	}

	return &models.Playlist{
		Name:          rootName,
		Path:          path,
		PlaylistNames: &playlists,
		Tracks:        &tracks,
	}
}

func (a *App) Play(track models.Track) {
	runtime.EventsEmit(a.ctx, "play:track", track)
}
