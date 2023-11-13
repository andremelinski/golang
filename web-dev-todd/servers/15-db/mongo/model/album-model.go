package model

type IAlbumProps struct {
	Title  string  `bson:"title" json:"title" valid:"notnull"`
	Artist string  `bson:"artist" json:"artist" valid:"notnull"`
	Price  float32 `bson:"price" json:"price" valid:"notnull"`
}
type IAlbumDB struct {
	ID          interface{} `bson:"_id" json:"id"`
	IAlbumProps `bson:",inline" json:"album_props"`
}