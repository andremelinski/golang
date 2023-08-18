package main

import "fmt"
//go run main.go > index.html will create a index.html file with the result fo this program
func main() {
	name := "Andre Melinski"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}