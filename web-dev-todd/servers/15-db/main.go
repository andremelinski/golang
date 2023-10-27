package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
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
    GetAllAlbums()
    GetAlbumById()
    UpdateAlbum()
    DeleteAlbum()
}

var db *sql.DB
var err error

var albumDB IAlbumDB

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
    defer db.Close()
    
    HandleRequests()    
}

func HandleRequests(){
    mux := httprouter.New()
	mux.POST("/", CreateAlbum)
	mux.GET("/", GetAllAlbums)
	mux.GET("/album/:id", GetAlbumById)
    mux.PUT("/album/:id", UpdateAlbum)
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

   newAlbumId, err := albumDB.DbCreateAlbm(albumData)
   if(err != nil) {
        http.Error(res, err.Error(), 500)
        return 
    }

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(newAlbumId)
}

func GetAllAlbums(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
    albums, err := albumDB.DbLoadAllAlbums()
    if(err != nil) {
        http.Error(res, err.Error(), 500)
        return 
    }
    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(albums)
}

func GetAlbumById(res http.ResponseWriter, _ *http.Request, param httprouter.Params){
    album, err := albumDB.DbLoadAlbumById(param.ByName("id"))
    if(err != nil) {
        http.Error(res, err.Error(), 500)
        return 
    }
    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(album)
}

func UpdateAlbum(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
    albumId := param.ByName("id")
    reqBody, _ := io.ReadAll(req.Body)

	albumData := IAlbumProps{}
	if err := json.Unmarshal(reqBody, &albumData); err != nil {
        http.Error(res, err.Error(), 500)
		return 
    }

   albumUpdated, err := albumDB.DbUpdateAlbum(albumData, albumId)
   if(err != nil) {
        http.Error(res, err.Error(), 500)
        return 
    }

    res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(albumUpdated)
}

func (albumDB IAlbumDB)DbCreateAlbm(albumData IAlbumProps)(int64, error){
     newAlbum, err := db.Exec("INSERT INTO album(title, artist, price) VALUES (?,?,?)", 
    albumData.Title, albumData.Artist, albumData.Price)

    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }

    id, err := newAlbum.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("LastInsertId in addAlbum: %v", err)
    }
    return id, nil
}

func(albumDb IAlbumDB) DbLoadAllAlbums()(*[]IAlbumDB, error){
    albums := &[]IAlbumDB{}

    rows, err := db.Query("SELECT * FROM album")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        alb := IAlbumDB{}
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, err
        }
        *albums = append(*albums, alb)
    }

    return albums, nil
}

func(albumDb IAlbumDB) DbLoadAlbumById(id string)(*IAlbumDB, error){
    row := db.QueryRowContext(context.TODO(), "SELECT * FROM album WHERE id=?", id)
    err := row.Scan(&albumDb.ID, &albumDb.IAlbumProps.Title, &albumDb.IAlbumProps.Artist, &albumDb.IAlbumProps.Price)
    if err != nil {
        return nil, err
    }
    return &albumDb, nil
}

func(albumDb IAlbumDB) DbUpdateAlbum(albumData IAlbumProps, id string)(int64, error){
    fieldsArr := reflect.TypeOf(albumData)
    valuesArr := reflect.ValueOf(albumData)
    fieldsNum := fieldsArr.NumField()
    updateConditition:= "UPDATE album SET "

    for i := 0; i < fieldsNum; i++ {
        field := fieldsArr.Field(i)
        fieldName := strings.ToLower(field.Name)
        value := valuesArr.Field(i)
        switch field.Type.Kind(){
        case reflect.String:
            if value.String() != "" {
                stringCondition := fieldName+"="+"'"+value.String()+"', "
                updateConditition += stringCondition
            }
    default:
        floatValue := value.Float()
        if floatValue > 0 {
                stringCondition := fieldName+"="+ fmt.Sprintf("%.2f", floatValue)+", "
                // stringCondition := fieldName+"="+ strconv.FormatFloat(floatValue, 'f', -1, 64)
                updateConditition += stringCondition
            }
        }
    }
    updateConditition = updateConditition[:len(updateConditition)-2] + " WHERE id="+id+";"
    // updateConditition = strings.Replace(updateConditition, ", WHERE", " WHERE", -1)

    albumUpdated, err := db.Exec(updateConditition)

        if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }

	row, err := albumUpdated.RowsAffected()

    if err != nil {
            return 0, fmt.Errorf("RowsAffected in updateAlbum: %v", err)
        }
    return row, nil
    
}

func validator[T any](dataObj T)error{
	_, err := govalidator.ValidateStruct(dataObj)
	if err!=nil{
		return err
	 }
	 return nil
}