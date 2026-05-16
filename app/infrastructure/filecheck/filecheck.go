package filecheck

import (
	"os"
	"path/filepath"
	"strings"
)

func IsEntryTrack(entry os.DirEntry) bool {
	name := entry.Name()
	ext := strings.ToLower(filepath.Ext(name))

	if name[0] == '.' {
		return false
	}

	formats := map[string]struct{}{
		"mp3":  {},
		"opus": {},
		"ogg":  {},
		"flac": {},
	}

	_, ok := formats[ext]
	return ok
}

func IsEntryPlaylist(entry os.DirEntry) bool {
	return entry.IsDir()
}
