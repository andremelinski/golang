package db

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// https://www.codingninjas.com/studio/library/promoted-methods-in-structure-in-go
type IUserProps struct {
	Name  string `json:"name" valid:"notnull"`
	Email  string `json:"email" valid:"notnull,email"`
	Password string `json:"password" valid:"notnull"`
	Age   int `json:"age" valid:"notnull"`
	Smoke bool `json:"smoke" valid:"-"`
}

type IUserLogin struct {
	Name  string `json:"name" valid:"notnull"`
	Password string `json:"password" valid:"notnull"`
}

type IUserPropsDB struct {
	Id    string 
	IUserProps
}
var userArr []IUserPropsDB

// var checkerKeysObj = map[string][]string{
// 	"user": {"Name", "Password","Age", "Smoke"},
// }

func init(){
	// https://hackwild.com/article/go-input-validation-and-testing/
	govalidator.SetFieldsRequiredByDefault(true)
}

func CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	reqBody, _ := io.ReadAll(req.Body)

	userData := IUserProps{}
	if err := json.Unmarshal(reqBody, &userData); err != nil {
        http.Error(res, err.Error(), 500)
		return
    }

	isValid := validator[IUserProps](userData)
	if(isValid!=nil){
		http.Error(res, isValid.Error(), 500)
		return
	}

	user := findUSerByName(userData.Name)
	userId := user.Id
	if user.Id != "" {
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode("user already exist")
		return
	}
	
	newUser, err := appendUser(userData)
	if err!=nil {
		http.Error(res, isValid.Error(), 500)
		return
	}
	userId = newUser.Id
	http.SetCookie(res, &http.Cookie{
		Name: "session-id",
		Value: userId,
		MaxAge: 10,
	})

	http.Redirect(res, req, "/login", 300)
}

func SignUp(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userFound:= alreadyLoggedIn(req)

	if(userFound.Id != ""){
		http.Redirect(res, req, "/login", 300)
		return
	}

	reqBody, _ := io.ReadAll(req.Body)

	userData := IUserLogin{}
	if err := json.Unmarshal(reqBody, &userData); err != nil {
        http.Error(res, err.Error(), 500)
		return
    }

	isValid := validator[IUserLogin](userData)
	if(isValid!=nil){
		http.Error(res, isValid.Error(), 500)
		return
	}

	userFound = findUSerByName(userData.Name)
	
	err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userData.Password))
	if err != nil {
		http.Error(res, "Username and/or password do not match", http.StatusForbidden)
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name: "session-id",
		Value: userFound.Id,
		MaxAge: 10,
	})

	http.Redirect(res, req, "/login", 300)
}

func findUSerByName(name string) *IUserPropsDB  {
	foundUser := IUserPropsDB{}
	for _, user := range userArr {
		if user.Name == name {
			foundUser = user
			break
		}
	}
	return &foundUser
}

func findUSerByCookieId(id string) *IUserPropsDB  {
	foundUser := IUserPropsDB{}
	for _, user := range userArr {
		if user.Id == id {
			foundUser = user
			break
		}
	}
	return &foundUser
}

func appendUser(userData IUserProps) (*IUserPropsDB, error){
	hashedPassword, err :=  bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost); 
	if err != nil {
		return nil, err
	}
	userData.Password = string(hashedPassword)
	newUser := IUserPropsDB{uuid.NewString(), userData}
	userArr = append(userArr, newUser)	
	return &newUser, nil
}

func alreadyLoggedIn(req *http.Request) *IUserPropsDB {
	c, err := req.Cookie("session")
	if err != nil {
		return nil
	}
	user := findUSerByCookieId(c.Value)
	return user
}

func validator[T any](dataObj T)error{
	_, err := govalidator.ValidateStruct(dataObj)
	if err!=nil{
		return err
	 }
	 return nil
}