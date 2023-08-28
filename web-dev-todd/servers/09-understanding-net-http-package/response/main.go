package main

import (
	"fmt"
	"net/http"
)

type number int

var d number

func (m number) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Mcleod-Key", "this is from mcleod")
	// res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader((http.StatusBadGateway))
	res.Write([]byte(`{message: "hello from json application"}`))
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}

func main() {
	http.ListenAndServe(":8080", d)
}
