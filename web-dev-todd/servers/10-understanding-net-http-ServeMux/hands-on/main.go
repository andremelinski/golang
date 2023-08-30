package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

// https://tutorialedge.net/golang/creating-restful-api-with-golang/
type IArticleProps struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	PublishedYear int    `json:"year"`
}

var pArticle *[]IArticleProps

func handleRequests() {
	mux := httprouter.New()
	mux.GET("/", homePage)
	mux.GET("/article", getAllArticles)
	mux.GET("/article/:id", getArticleById)
	mux.POST("/article", createArticle)
	mux.PUT("/article/:id", updateArticleById)
	mux.DELETE("/article/:id", deleteArticleById)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func main() {
	fmt.Println("Creating Articles")
	Articles := []IArticleProps{
		{Id: "1", Title: "Hello", Desc: "Article Description 1", Content: "Article Content 1", PublishedYear: 2020},
		{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2", PublishedYear: 2023},
		{Id: "3", Title: "Hello 3", Desc: "Article Description 3", Content: "Article Content 3", PublishedYear: 2023},
	}
	pArticle = &Articles
	fmt.Println("Articles Created")
	handleRequests()
}

func homePage(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(res).Encode("Hello from the main page")
}

func getAllArticles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(*pArticle)
	return
}

func getArticleById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	article := findArticeById(id)

	if article.Id == "" {
		res.WriteHeader(http.StatusNoContent)
		json.NewEncoder(res).Encode("{}")
	} else {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(article)
	}

	return
}

func createArticle(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(req.Body)

	dataArticle := IArticleProps{}

	err := json.Unmarshal(reqBody, &dataArticle)

	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	*pArticle = append(*pArticle, dataArticle)

	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%+v \n", string(reqBody))
	json.NewEncoder(res).Encode(*pArticle)

}

func getArticleByYear(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	params.ByName("year")
	fmt.Println("Endpoint Hit: getArticleById")
}

func updateArticleById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	article := findArticeById(id)

	if article.Id == "" {
		res.WriteHeader(http.StatusNoContent)
		json.NewEncoder(res).Encode("{}")
	} else {
		reqBody, _ := io.ReadAll(req.Body)
		dataArticle := IArticleProps{}

		err := json.Unmarshal(reqBody, &dataArticle)

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		// https://www.golangprograms.com/how-to-get-struct-variable-information-using-reflect-package.html
		//In Go, it's not possible to partially update a struct dynamically in the same way you might do with some dynamic languages.
		//When you want to update a struct's field, you generally have to assign a new value to the entire field
		e := reflect.ValueOf(&dataArticle).Elem()
		updateArticle := updateFields(&article, e)
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(updateArticle)
	}

	return
}

func deleteArticleById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	res.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	articles := *pArticle
	for index, article := range articles {
		if article.Id == id {
			// updates our Articles array to remove the
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
	*pArticle = articles

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(*pArticle)
	return
}

func updateFields(article *IArticleProps, e reflect.Value) *IArticleProps {
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name

		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v \n", varName, varValue)
		switch varName {
		case "Desc":
			newDesc := varValue.(string)
			if len(newDesc) > 0 {
				article.Desc = newDesc
			}
		case "Title":
			newTitle := varValue.(string)
			if len(newTitle) > 0 {
				article.Title = newTitle
			}
		case "Content":
			newContent := varValue.(string)
			if len(newContent) > 0 {
				article.Content = newContent
			}
		case "PublishedYear":
			newPublishedYear := varValue.(int)
			if newPublishedYear > 0 {
				article.PublishedYear = varValue.(int)
			}

		}
	}
	return article
}

func findArticeById(id string) IArticleProps {
	foundArticle := IArticleProps{}
	for _, article := range *pArticle {
		if article.Id == id {

			foundArticle = article
		}
	}
	return foundArticle
}
