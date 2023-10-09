package db

import (
	"encoding/json"
	"fmt"
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
	// if(!isValid){
	// 	res.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(res).Encode("Keys alowed: "+ strings.Join(checkerKeysObj["user"]," "))
	// 	return
	// }

	user := findUSerByName(userData.Name)
	fmt.Println(user)
	userId := user.Id
	if user.Id == "" {
		newUser, err := appendUser(userData)
		if err!=nil {
			http.Error(res, isValid.Error(), 500)
			return
		}
		fmt.Println(newUser)
		userId = newUser.Id

	}
	http.SetCookie(res, &http.Cookie{
		Name: "session-id",
		Value: userId,
		MaxAge: 10,
	})

	// res.WriteHeader(http.StatusCreated)
	// json.NewEncoder(res).Encode(userData)
	// http.Redirect(res, req, "/login", 300)
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

func validator[T IUserProps](dataObj T)error{
	// v := reflect.ValueOf(	dataObj)
	// keysToCheck := checkerKeysObj[keyToCheck]

	// return v.NumField() == len(keysToCheck)
	_, err := govalidator.ValidateStruct(dataObj)
	if err!=nil{
		return err
	 }
	 return nil
}