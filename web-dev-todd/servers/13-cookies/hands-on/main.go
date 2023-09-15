package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func handleRequests() {
	mux := httprouter.New()
	mux.GET("/", homePage)
	mux.GET("/getcookie", getCookie)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	handleRequests()
}

func homePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie, err := req.Cookie("counter-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter-cookie",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi((cookie.Value))
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(res, cookie)
	json.NewEncoder(res).Encode("cookie saved")
}

func getCookie(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("counter-cookie")
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	json.NewEncoder(res).Encode(c)

}
