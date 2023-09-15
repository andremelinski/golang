package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleRequests() {
	mux := httprouter.New()
	mux.GET("/", homePage)
	mux.GET("/getcookie", getCookie)
	mux.GET("/delCookie", delCookie)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	handleRequests()
}

func homePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	a := &http.Cookie{
		Name:  "first-cookie",
		Value: "cookie-value",
	}
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a)
	http.SetCookie(res, a)
	json.NewEncoder(res).Encode("cookie saved")
}

func getCookie(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("first-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	json.NewEncoder(res).Encode(c)

}

func delCookie(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("first-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	c.MaxAge = 0
	http.SetCookie(res, c)
	json.NewEncoder(res).Encode("cookie deleted")

}
