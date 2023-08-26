package main

import (
	"fmt"
	"log"
	"net"
)

// run "02_read-scanner"
// run "dial-write"
// deal will create a connection to tcp server writing "I dialed you." and read scanner will print this and close connection which will print "code got here"
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
