package data

//https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
import (
	"context"
	"reflect"
	"strings"

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

func(albumDb DbAlbumRepo) DbUpdateByIdAlbum(id string, updateInfo *model.IAlbumProps)(*model.IAlbumProps, error){
	collection := albumDb.initCollection()
	album := model.IAlbumProps{}

	ojbID, _ := primitive.ObjectIDFromHex(id)
	filterOptions := bson.D{{"_id", ojbID}}

	toUpdate := prepareDateToUpdate(*updateInfo)

	if len(toUpdate)>0 {
		err := collection.FindOneAndUpdate(
				context.TODO(),
				filterOptions,
				bson.D{{"$set", toUpdate}},
			).Decode(&album)
	
			//   fmt.Println(u)
		if(err!= nil){
			return nil, err
		}
	}

	return &album, nil

}

func(albumDb DbAlbumRepo) DbDeleteByIdAlbum(id string)(*model.IAlbumDB, error){
	collection := albumDb.initCollection()
	albumDeleted := model.IAlbumDB{}
	ojbID, _ := primitive.ObjectIDFromHex(id)
	filterOptions := bson.M{"_id": ojbID}

	err := collection.FindOneAndDelete(context.TODO(), filterOptions).Decode(&albumDeleted)
	if(err!= nil){
		return nil, err
	}

	return &albumDeleted, nil
}

func prepareDateToUpdate(updateInfo model.IAlbumProps) []bson.E{
	update := []bson.E{}
	fieldsArr := reflect.TypeOf(updateInfo)
    valuesArr := reflect.ValueOf(updateInfo)
    fieldsNum := fieldsArr.NumField()

	for i := 0; i < fieldsNum; i++ {
		field := fieldsArr.Field(i)
		fieldName := strings.ToLower(field.Name)
		fieldValue := valuesArr.Field(i)
		switch field.Type.Kind(){
			case reflect.String:
				if fieldValue.String() != "" {
					update = append(update, bson.E{fieldName, fieldValue.String()})
				}
			default:
				floatValue := fieldValue.Float()
				if floatValue > 0 {
						update = append(update, bson.E{fieldName, floatValue})
					}
        }
	}
	return update
}