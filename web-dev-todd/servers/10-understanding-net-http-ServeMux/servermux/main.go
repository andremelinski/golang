package main

import (
	"io"
	"net/http"
)

// type r1 int
// type r2 int

// var dataFromR1 r1
// var dataFromR2 r2

// func (d r1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "route 1 called")
// }

// func (d2 r2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "route 2 called")
// }

func dataFromR2(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route 1 called")
}

func dataFromR1(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route 2 called")
}

func main() {
	// mux := http.NewServeMux()
	// mux.Handle("/r1", dataFromR1)
	// mux.Handle("/r2", dataFromR2)
	// http.ListenAndServe(":8080", mux)

	// OR
	// http.Handle("/r1", dataFromR1)
	// http.Handle("/r2", dataFromR2)
	// http.ListenAndServe(":8080", nil)

	// OR
	// convert dataFrom1 to a handler func and then it has that serve Http method with response writer and pointer to a request.
	// http.Handle("/r1", http.HandlerFunc(dataFromR1))
	// http.Handle("/r2", http.HandlerFunc(dataFromR2))
	// http.ListenAndServe(":8080", nil)

	// OR (most commun)
	http.HandleFunc("/r1", dataFromR1)
	http.HandleFunc("/r2", dataFromR2)
	http.ListenAndServe(":8080", nil)
}
