package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleRequests() {
	mux := httprouter.New()
	mux.GET("/", homePage)
	mux.GET("/redirect", redirectRouter)
	mux.GET("/redirect2", redirectRouter2)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	handleRequests()

}

func homePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(res).Encode("Hello from the main page 2")
}

// internal redirect
func redirectRouter(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(res, req, "/", 300)
}

// outside redirect
func redirectRouter2(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Redirect(res, req, "http://www.google.com", http.StatusPermanentRedirect)
}
