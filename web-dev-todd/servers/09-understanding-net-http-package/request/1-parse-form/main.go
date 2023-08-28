package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type number int

var i number
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func (m number) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(req.Form)

	// req.Form is for values from both the form and the URL.
	// req.PostForm is just from the form
	//you got to run parse form on your request before you get those values
	tpl.ExecuteTemplate(res, "index.html", req.Form)
}
func main() {
	http.ListenAndServe(":8080", i)
}
