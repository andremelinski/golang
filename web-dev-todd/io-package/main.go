package main

import (
	"fmt"
	"io"
	"os"
)

// go run main.go > index.html will create a index.html file with the result fo this program
func main() {
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	count, err := writer.Write([]byte("Hello"))
	fmt.Println(count, err)

	countByString, err2ByString := io.WriteString(writer, " Word")
	fmt.Println(countByString, err2ByString)

	//***** OPEN FILE *****
	fileOpen, errOPen := os.Open("file.txt")
	fmt.Println("Open", fileOpen, errOPen)

	// https://linuxhint.com/golang-os-open/
	// read or create if not exists
	fileOPenFile, errOPenFile := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Println("OPenFile", fileOPenFile, errOPenFile)

	//***** READING FILE *****
	reader := io.Reader(fileOpen)
	// read 1024 bytes at a time.
	buffer := make([]byte, 1024)
	n, err := reader.Read(buffer)
	fmt.Printf("reader.Read: Read n={%v}, err={%v},buffer={%v}\n", n, err, string(buffer))

	contentReadAll, errReadAll := io.ReadAll(reader)
	fmt.Printf("ReadAll: err={%v}, buffer={%v}\n", errReadAll, string(contentReadAll))

	// https://gosamples.dev/read-file/
	content, err := os.ReadFile("file.txt")
	fmt.Printf("ReadFile: err={%v},buffer={%v}\n", err, string(content))

	//***** SEEKING IN A FILE *****
	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)
	seeker.Seek(0, io.SeekCurrent)
	seeker.Seek(-5, io.SeekEnd)
	content2ReadAll, errReadAll2 := io.ReadAll(reader)
	fmt.Printf("ReadAll: err={%v}, buffer={%v}\n", errReadAll2, string(content2ReadAll))
	file.Close()
}
