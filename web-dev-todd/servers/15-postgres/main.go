package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-postgres/controller"
	"github.com/andremelinski/web-dev-todd/servers/15-postgres/infra"
	"github.com/julienschmidt/httprouter"
)

func main() {
	sqlDb, err := infra.InitDataBaseConnection().ConnectMySQLDB()

	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Connected!")

	// repos := repository.InitRepositories(sqlDb)
	// repos.CompanyRepo
	mux := httprouter.New()

	// get album controllers
	controllers := controller.InitControllers(sqlDb)

	// Company
	mux.POST("/", middlewareContentType(controllers.TablesDb.CreateDbTables))
	mux.POST("/company", middlewareContentType(controllers.CompanyController.CreateCompanyController))
	// mux.GET("/company", middlewareContentType(controllers.companyController.GetCompanys))
	// mux.GET("/company/:id", middlewareContentType(controllers.companyController.GetCompanyById))
	// mux.PUT("/company/:id", middlewareContentType(controllers.companyController.UpdateByIdCompany))
	// mux.DELETE("/company/:id", middlewareContentType(controllers.companyController.DeleteByIdCompany))
	// // Employee
	// mux.POST("/employee", middlewareContentType(albumController.CreateEmployee))
	// mux.GET("/employee", middlewareContentType(albumController.GetEmployees))
	// mux.GET("/employee/:id", middlewareContentType(albumController.GetEmployeeById))
	// mux.PUT("/employee/:id", middlewareContentType(albumController.UpdateByIdEmployee))
	// mux.DELETE("/employee/:id", middlewareContentType(albumController.DeleteByIdEmployee))
	log.Fatal(http.ListenAndServe(":8080", mux))

	if err != nil {
		log.Fatalln(err)
	}

}

func middlewareContentType(next httprouter.Handle) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
		res.Header().Set("Content-Type", "application/json")
		if next != nil {
			next(res, req, p)
		}
	}
}