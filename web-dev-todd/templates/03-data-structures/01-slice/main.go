package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("slice.html", "slice-with-index.html"))
}

func main() {
	si := []int{01, 2, 3, 4}
	ss := []string{"A", "B", "C", "D"}
	err := tpl.ExecuteTemplate(os.Stdout, "slice.html", si)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "slice-with-index.html", ss)
	if err != nil {
		log.Fatalln(err)
	}
}
