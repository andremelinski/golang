package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.html"))
}

// https://gobyexample.com/time
// https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go
func monthDayYear(t time.Time) string {
	aqui := t.Format(time.Kitchen)
	fmt.Println(aqui)
	// MM-DD-YYYY
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
