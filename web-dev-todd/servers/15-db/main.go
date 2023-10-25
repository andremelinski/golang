package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)
type IAlbumProps struct {
    Title  string `json:"title" valid:"notnull"`
    Artist string `json:"artist" valid:"notnull"`
    Price  float32 `json:"price" valid:"notnull"`
}
type IAlbumDB struct{
    ID     int64
    IAlbumProps
}

type IAlbumRepository interface{
    CreateAlbum()
    GetAllAlbuns()
    GetAlbumById()
    UpdateAlbum()
    DeleteAlbum()
}

var db *sql.DB
var err error

func init(){
	// https://hackwild.com/article/go-input-validation-and-testing/
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
 // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "mypassword",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testdb",
		AllowNativePasswords: true,
    }

	db, err = sql.Open("mysql", cfg.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 1)
	if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    HandleRequests()    
}

func HandleRequests(){
    mux := httprouter.New()
	mux.POST("/", CreateAlbum)
	// mux.GET("/", GetAllAlbuns)
	// mux.GET("/album/:id", GetAlbumById)
    // mux.PUT("/album/:id", UpdateAlbum)
	// mux.DELETE("/album/:id", DeleteAlbum)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func CreateAlbum(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    reqBody, _ := io.ReadAll(req.Body)

	albumData := IAlbumProps{}
	if err := json.Unmarshal(reqBody, &albumData); err != nil {
        http.Error(res, err.Error(), 500)
		return 
    }
    isValid := validator[IAlbumProps](albumData)
    if(isValid!=nil){
		http.Error(res, isValid.Error(), 500)
		return 
	}

   newAlbumId, err := albumData.DbCreateAlbm()
   if(err != nil) {
        http.Error(res, err.Error(), 500)
        return 
    }

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(newAlbumId)

}

func (albumData IAlbumProps)DbCreateAlbm()(int64, error){
     newAlbum, err := db.Exec("INSERT INTO album(title, artist, price) VALUES (?,?,?)", 
    albumData.Title, albumData.Artist, albumData.Price)

    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }

    id, err := newAlbum.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}

func validator[T any](dataObj T)error{
	_, err := govalidator.ValidateStruct(dataObj)
	if err!=nil{
		return err
	 }
	 return nil
}