package main

import (
	"fmt"
	"net/http"
)

// any value of type hot dog is implicitly implementing the handler interface.
type hotdog int

// "Handler"
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
