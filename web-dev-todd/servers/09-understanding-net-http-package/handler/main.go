package main

import (
	"fmt"
	"net/http"
)

type hotdog int

var test hotdog

// type hotdog implements a ServerHttp functin and handle the HTTP connection
func (x hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf(req.RequestURI)
	fmt.Println("Any code you want in this func \n")
}

// start the server
func main() {
	http.ListenAndServe(":8080", test)
}
