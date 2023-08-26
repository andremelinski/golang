package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// to read from a connection u need a go function
// go run main and at cmd telnet localhost 8080. Every time you write on the cmd, it will appear
// on the vsCode terminal
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}
func handle(connection net.Conn) {
	setTimeOut := connection.SetDeadline(time.Now().Add(10 * time.Second))

	if setTimeOut != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(connection)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// will scan and it'll continue to give me a line each time it scans.
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(connection, "I heard you say: %s\n", ln)
	}
	defer connection.Close()

	// we never get here
	// we have an open stream http connection, scanner.Scan will never close
	// after 10s connection is closed and code got here appears
	fmt.Println("Code got here.")
}
