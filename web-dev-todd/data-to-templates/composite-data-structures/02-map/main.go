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
	tpl = template.Must(template.ParseFiles("map.html"))
}

func main() {
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "map.html", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
