package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// run "01_write"
// run "06_dial-read"
// deal will create a connection to tcp server and read info from it, returning what's written on the program
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
