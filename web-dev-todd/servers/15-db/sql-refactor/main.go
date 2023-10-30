package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	album_controller "github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/controller/album"
	"github.com/andremelinski/web-dev-todd/servers/15-db/sql-refactor/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg := mysql.Config{
        User:   "root",
        Passwd: "mypassword",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testdb",
		AllowNativePasswords: true,
    }

	db, err := sql.Open("mysql", cfg.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 1)
	if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	HandleRequests(db)
}

func HandleRequests(db *sql.DB) {

	repos := repository.InitRepositories(db)

	mux := httprouter.New()

	// controllers := controller.InitControllers(repos)

	// get album controllers
	albumController := album_controller.InitAlbumController(repos.AlbumRepo)
	
	mux.POST("/", middlewareContentType(albumController.CreateAlbum))
	// mux.GET("/", middlewareContentType(GetAllAlbums))
	// mux.GET("/album/:id", middlewareContentType(GetAlbumById))
	// mux.PUT("/album/:id", middlewareContentType(UpdateAlbum))
	// mux.DELETE("/album/:id", middlewareContentType(DeleteAlbum))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func middlewareContentType(next httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		res.Header().Set("Content-Type", "application/json")
		if next != nil {
			next(res, req, p)
		}
	}
}