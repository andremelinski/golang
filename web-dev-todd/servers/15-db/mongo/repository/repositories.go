package repository

import (
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/repository/data"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	AlbumRepo *data.DbAlbumRepo
}

func InitRepositories(db *mongo.Database)*Repositories{
	albumRepo := data.NewAlbumRepo(db)
	return &Repositories{
		albumRepo,
	}
}