package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type IUserProps struct {
	Name  string `json:"name"`
	Age   int `json:"age"`
	Smoke bool `json:"smoke"`
}

type IUserPropsDB struct {
	Id    string
	Name  string 
	Age   int 
	Smoke bool 
}
var userArr []IUserPropsDB

var checkerKeysObj = map[string][]string{
		"user": {"Name", "Age", "Smoke"},
	}

func CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqBody, _ := io.ReadAll(req.Body)

	userData := IUserProps{}
	if err := json.Unmarshal(reqBody, &userData); err != nil {
        http.Error(res, err.Error(), 500)
		return
    }

	isValid := validator[IUserProps](userData, "user")

	if(!isValid){
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode("Keys alowed: "+ strings.Join(checkerKeysObj["user"]," "))
		return
	}

	user := findUSerByName(userData.Name)
	if user.Id == "" {
		user = appendUser(userData)
	}
	http.SetCookie(res, &http.Cookie{
		Name: "session-id",
		Value: user.Id,
		MaxAge: 10,
	})

	// res.WriteHeader(http.StatusCreated)
	// json.NewEncoder(res).Encode(userData)
	http.Redirect(res, req, "/login", 300)
}

func SignUp(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie, err := req.Cookie("session-id")
	if err == http.ErrNoCookie {
		//  http.Error(res, err.Error(), 500)
		fmt.Println( err.Error())
		http.Redirect(res, req, "/logout", 300)
		return
	}
	userFound := findUSerByCookieId(cookie.Value)

	if(userFound.Id == ""){
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("user not Found")
		return
	}
	// res.WriteHeader(http.StatusCreated)
	// json.NewEncoder(res).Encode(userFound)
	http.Redirect(res, req, "/login", 300)
}

func findUSerByName(name string) IUserPropsDB  {
	foundUser := IUserPropsDB{}
	for _, user := range userArr {
		if user.Name == name {
			foundUser = user
			break
		}
	}
	return foundUser
}

func findUSerByCookieId(id string) IUserPropsDB  {
	foundUser := IUserPropsDB{}
	for _, user := range userArr {
		if user.Id == id {
			foundUser = user
			break
		}
	}
	return foundUser
}

func appendUser(userData IUserProps) IUserPropsDB{
	newUser := IUserPropsDB{uuid.NewString(), userData.Name, userData.Age, userData.Smoke}
	userArr = append(userArr, newUser)	
	return newUser
}

func validator[T any](dataObj T, keyToCheck string)bool{
	v := reflect.ValueOf(	dataObj)
	keysToCheck := checkerKeysObj[keyToCheck]

	if (v.NumField()!= len(keysToCheck)){
		return false
	}
	return true
}