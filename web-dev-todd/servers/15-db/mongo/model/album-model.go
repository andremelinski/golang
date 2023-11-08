package model

type IAlbumProps struct {
	Title  string  `json:"title" valid:"notnull"`
	Artist string  `json:"artist" valid:"notnull"`
	Price  float32 `json:"price" valid:"notnull"`
}
type IAlbumDB struct {
	ID interface{}
	IAlbumProps
}