package interfaces

import "github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"

type IAlbumRepositoryInterface interface {
	DbCreateAlbum(albumData model.IAlbumProps) (interface{}, error)
	DbReadAll() (*[]model.IAlbumDB, error)
	DbreadByIdAlbum(id string) (*model.IAlbumProps, error)
	DbUpdateByIdAlbum(id string, albumData *model.IAlbumProps, ) (*model.IAlbumProps, error)
	DbDeleteByIdAlbum(id string) (*model.IAlbumDB, error)
}
