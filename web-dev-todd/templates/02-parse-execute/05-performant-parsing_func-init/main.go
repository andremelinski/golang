package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// init func is called once when program runs
func init() {
	// Retuns a pointer where those files are storaged
	// Must do the error checking
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	err := tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
