package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitSQLiteStorage(path string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(Query); err != nil {
		return nil, err
	}

	return db, nil
}
