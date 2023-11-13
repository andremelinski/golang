package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/controller"
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/infra"
	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/repository"
	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
)

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	mongodb, cancel, err := infra.InitDataBaseConnection().ConnectMongoDB()
	
	defer cancel()
	  if(err != nil){
		  log.Fatalln(err)
		  return 
	  }
	  fmt.Println("Connected!")
	
	  repos := repository.InitRepositories(mongodb)

	mux := httprouter.New()

	// get album controllers
	albumController := controller.InitAlbumController(repos.AlbumRepo)
	
	mux.POST("/", middlewareContentType(albumController.CreateAlbum))
	mux.GET("/", middlewareContentType(albumController.GetAlbums))
	mux.GET("/:id", middlewareContentType(albumController.GetAlbumById))
	mux.PUT("/:id", middlewareContentType(albumController.UpdateByIdAlbum))
	mux.DELETE("/:id", middlewareContentType(albumController.DeleteByIdAlbum))
	log.Fatal(http.ListenAndServe(":8080", mux))

	if(err != nil){
		log.Fatalln(err)
	}

}

func middlewareContentType(next httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		res.Header().Set("Content-Type", "application/json")
		if next != nil {
			next(res, req, p)
		}
	}
}