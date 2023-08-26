package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// This creates a new template and parses all files that is passed out.
	// Retuns a pointer where those files are storaged
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatal(err)
	}

	files := []string{"one.gmao", "two.gmao", "vespa.gmao"}
	// adding those 3 files on the pointer
	tpl, err = tpl.ParseFiles(files...)
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
