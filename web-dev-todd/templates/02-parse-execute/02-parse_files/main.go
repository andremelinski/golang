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
	// Execute it to allow us to display dynamic data
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	// create a new file
	nf, err := os.Create("index2.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()

	// send the data from tpl to nf
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
