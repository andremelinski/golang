package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

const PORT = 8080

var errNew string
var http_status int

func handleRequests() {
	mux := httprouter.New()
	mux.GET("/", homePage)
	// mux.GET("/photo", getAllPhotos)
	// mux.GET("/photo/:id", getPhotoById)
	mux.POST("/photo", createPhoto)
	// mux.DELETE("/photo/:id", deletePhotoById)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	handleRequests()

}

func homePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(res).Encode("Hello from the main page")
}

// https://medium.com/@owlwalks/dont-parse-everything-from-client-multipart-post-golang-9280d23cd4ad
func createPhoto(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// 32 MB is the default used by FormFile() function
	if err := req.ParseMultipartForm(32 << 20); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// get files from file variable in multi-part/form
	files := req.MultipartForm.File["file"]

	for _, fileHeader := range files {

		file, err := fileHeader.Open()
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}
		defer file.Close()

		buffer := make([]byte, 1024)

		_, err = file.Read(buffer)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}

		mime := checkMimeType(buffer)

		if len(mime) == 0 {
			break
		}
		// https://freshman.tech/snippets/go/reset-file-pointer/
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusInternalServerError
			break
		}

		//creating empty file with final file name, date now and same file extension
		fileName := strings.ReplaceAll(strings.Split(fileHeader.Filename, ".")[0], " ", "_")
		newFile, err := os.Create(fmt.Sprintf("./uploads/%v%d%s", fileName, time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusBadRequest
			break
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			errNew = err.Error()
			http_status = http.StatusBadRequest
			break
		}

	}
	message := "file uploaded successfully"
	if errNew != "" {
		message = errNew
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(message)
}

func checkMimeType(buffer []byte) string {
	filetype := http.DetectContentType(buffer)

	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" || filetype == "application/octet-stream" {
		errNew = "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image"
		http_status = http.StatusBadRequest
		return ""
	}
	return filetype
}
