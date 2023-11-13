package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/controller/validator"
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/interfaces"
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/model"
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/repository/data"
	"github.com/julienschmidt/httprouter"
)

type AlbumController struct {
	validate validator.Ivalidator
	service interfaces.IAlbumRepositoryInterface
}

func InitAlbumController(albumRepo *data.DbAlbumRepo) *AlbumController{
	return &AlbumController{
		service: albumRepo,
		validate: *validator.InitValidator(),
	}
}

func (albumControllerProps AlbumController) CreateAlbum(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqBody, _ := io.ReadAll(req.Body)

	albumData := model.IAlbumProps{}
	if err := json.Unmarshal(reqBody, &albumData); err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	isValid := albumControllerProps.validate.AlbumValidator(albumData)
    if(isValid!=nil){
		http.Error(res, isValid.Error(), 500)
		return 
	}

	id, err := albumControllerProps.service.DbCreateAlbum(albumData)
	if(err != nil){
		json.NewEncoder(res).Encode(err)
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(id)
}

func (albumControllerProps AlbumController) GetAlbums(res http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	albums, err := albumControllerProps.service.DbReadAll()

	if(err != nil){
		json.NewEncoder(res).Encode(err)
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(albums)
}
func (albumControllerProps AlbumController) GetAlbumById(res http.ResponseWriter, _ *http.Request, params httprouter.Params){
	album, err := albumControllerProps.service.DbreadByIdAlbum(params.ByName("id"))

	if(err != nil){
		json.NewEncoder(res).Encode(err)
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(album)
}