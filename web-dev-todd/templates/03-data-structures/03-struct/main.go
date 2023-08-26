package main

import (
	"log"
	"os"
	"text/template"
)

type IFile struct {
	Name        string
	LifeMeaning string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	dataFILe := IFile{"Andre", "money"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", dataFILe)
	if err != nil {
		log.Fatalln(err)
	}
}
