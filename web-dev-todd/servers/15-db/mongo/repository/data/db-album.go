package data

import (
	"context"

	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbAlbumRepo struct {
	db *mongo.Database
}

func NewAlbumRepo(db *mongo.Database)*DbAlbumRepo{
	return &DbAlbumRepo{
		db: db,
	}
}

func(albumDb DbAlbumRepo) DbCreateAlbum(albumData model.IAlbumProps)(interface{}, error){
	collection := albumDb.db.Collection("album")
	res, err := collection.InsertOne(context.Background(), albumData)

	if(err != nil){
		return nil, err
	}
	return res.InsertedID, nil
}