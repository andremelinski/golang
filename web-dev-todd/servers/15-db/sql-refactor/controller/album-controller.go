package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/data"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/interfaces"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/model"
	"github.com/julienschmidt/httprouter"
)

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type (
	// AlbumController represents the controller for operating on the Album resource
	AlbumController struct{
		serviceAlbum interfaces.IAlbumRepositoryInterface
	}
)

func InitAlbumControllers(albumRepo *data.AlbumRepo) *AlbumController {
	return &AlbumController{
		serviceAlbum: albumRepo,
	}
}

func (albumController AlbumController) CreateAlbum(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqBody, _ := io.ReadAll(req.Body)

	albumData := model.IAlbumProps{}
	if err := json.Unmarshal(reqBody, &albumData); err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	
	int, err := albumController.serviceAlbum.DbCreateAlbum(albumData)
	if(err != nil){
		json.NewEncoder(res).Encode(err)
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(int)

}
