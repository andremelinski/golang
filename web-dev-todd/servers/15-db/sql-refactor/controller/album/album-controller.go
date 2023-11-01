package album_controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/controller/album/validator"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/data"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/interfaces"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/model"
	"github.com/julienschmidt/httprouter"
)

// Apply Album interface by inheratence
type AlbumController struct {
	service interfaces.IAlbumRepositoryInterface
	validate validator.Ivalidator
}

func InitAlbumController(albumRepo *data.AlbumRepo) *AlbumController{
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

	int, err := albumControllerProps.service.DbCreateAlbum(albumData)
	if(err != nil){
		json.NewEncoder(res).Encode(err)
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(int)

}