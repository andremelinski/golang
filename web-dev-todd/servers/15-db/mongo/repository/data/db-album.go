package data

//https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
import (
	"context"

	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func(albumDb DbAlbumRepo) initCollection() *mongo.Collection {
	return albumDb.db.Collection("album")
}

func(albumDb DbAlbumRepo) DbCreateAlbum(albumData model.IAlbumProps)(interface{}, error){
	collection := albumDb.initCollection()
	res, err := collection.InsertOne(context.Background(), albumData)

	if(err != nil){
		return nil, err
	}
	return res.InsertedID, nil
}

func(albumDb DbAlbumRepo) DbreadByIdAlbum(id string)(*model.IAlbumProps, error){
	collection := albumDb.initCollection()
	
	album := model.IAlbumProps{}

	ojbID, _ := primitive.ObjectIDFromHex(id)
	filterOptions := bson.M{"_id": ojbID}

	err := collection.FindOne(context.Background(), filterOptions).Decode(&album)

	if(err != nil){
		return nil, err
	}

	return &album, nil
}

func(albumDb DbAlbumRepo) DbReadAll()(*[]model.IAlbumDB, error){
	collection := albumDb.initCollection()
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if(err != nil){
		return nil, err
	}
	
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	albums := []model.IAlbumDB{}
	//first way to load all 
	// if err = cursor.All(context.Background(), &albums); err != nil {
	// 	return nil, err
	// }
	
	// second way to load all 
	for cursor.Next(context.Background()) {
		// decode to struct 
		album := model.IAlbumDB{}
		err := cursor.Decode(&album)
		if(err != nil){
			return nil, err
		}
		albums = append(albums, album)
	}
	return &albums, nil
}