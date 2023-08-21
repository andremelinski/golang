package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// go run main.go > index.html will create a index.html file with the result fo this program
// go run main.go <Name> will add the <Name> in the file, else use default (Andre Melinski)
// and generate a copy from the final final as copyIndex
func main() {
	name := "Andre Melinski"
	if len(os.Args) > 1 {
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		name = os.Args[1]
	}

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

	// create empty file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// copy an empty file and add the string on it
	io.Copy(nf, strings.NewReader(tpl))

	// ***** FILE COPY *****
	sourceFile, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	copyFile, errSrc := os.Create("copyIndex.html")
	if errSrc != nil {
		log.Fatal(errSrc)
	}

	bytesCopied, _ := io.Copy(copyFile, sourceFile)
	defer copyFile.Close()
	log.Printf("Copied %d bytes.", bytesCopied)
}
