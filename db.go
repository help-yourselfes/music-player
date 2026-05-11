package main

import (
	"music-player/app/infrastructure/sqliteRepos"
	"music-player/app/infrastructure/storage"
	"music-player/app/services"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func (a *App) ConnectDB() error {
	basepath, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	name := "player.db"
	path := filepath.Join(basepath, name)

	db, err := storage.InitSQLiteStorage(path)
	if err != nil {
		return err
	}

	trackRepo := sqliteRepos.NewTrackRepoSQLite(db)
	playlistRepo := sqliteRepos.NewPlaylistRepoSQLite(db)

	appService := services.NewAppService(trackRepo, playlistRepo)
	a.service = *appService

	return nil
}
