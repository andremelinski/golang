package interfaces

import "github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"

type IAlbumRepositoryInterface interface {
	DbCreateAlbum(albumData model.IAlbumProps) (interface{}, error)
	// DbReadAll() (*[]model.IAlbumProps, error)
	// DbreadByIdAlbum(id string) (*model.IAlbumProps, error)
	// DbUpdateByIdAlbum(albumData model.IAlbumProps, id string) (int64, error)
	// DbDeleteByIdAlbum(id string) (int64, error)
}
