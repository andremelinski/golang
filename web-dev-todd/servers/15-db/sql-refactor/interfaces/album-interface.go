package interfaces

import "github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/model"

type IAlbumRepositoryInterface interface {
	DbCreateAlbum(albumData model.IAlbumProps) (int64, error)
	// DbReadAll() (*[]model.IAlbumProps, error)
	// DbreadByIdAlbum(id string) (*model.IAlbumProps, error)
	// DbUpdateByIdAlbum(albumData model.IAlbumProps, id string) (int64, error)
	// DbDeleteByIdAlbum(id string) (int64, error)
}
