package data

import (
	"database/sql"
	"fmt"

	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/model"
)
type AlbumRepo struct {
	db *sql.DB
}

func NewAlbumRepo(db *sql.DB) *AlbumRepo {
	return &AlbumRepo{
		db: db,
	}
}

func (albumDB *AlbumRepo) DbCreateAlbum(albumData model.IAlbumProps) (int64, error) {
	newAlbum, err := albumDB.db.Exec("INSERT INTO album(title, artist, price) VALUES (?,?,?)",
		albumData.Title, albumData.Artist, albumData.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := newAlbum.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("LastInsertId in addAlbum: %v", err)
	}
	return id, nil
}