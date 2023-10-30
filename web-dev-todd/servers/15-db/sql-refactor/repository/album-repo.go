package repository

import (
	"database/sql"

	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/data"
)

type Repositories struct {
	AlbumRepo *data.AlbumRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *sql.DB) *Repositories {
	albumRepo := data.NewAlbumRepo(db)
	return &Repositories{AlbumRepo: albumRepo}
}