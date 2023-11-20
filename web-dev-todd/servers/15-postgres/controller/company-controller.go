package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-postgres/domain/interfaces"
	"github.com/andremelinski/web-dev-todd/servers/15-postgres/repository/data"
	"github.com/julienschmidt/httprouter"
)

type ICompanyController struct {
	companyDb data.DbCompanyRepo
}

func InitCompanyController(companyRepo data.DbCompanyRepo)*ICompanyController{
	return &ICompanyController{
		companyRepo,
	}
}

func(companyController ICompanyController) CreateCompanyController(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	reqBody, _ := io.ReadAll(req.Body)
	companyInfo := interfaces.ICompanyProps{}
	
	if err := json.Unmarshal(reqBody, &companyInfo); err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	companyId, err := companyController.companyDb.CreateCompany(companyInfo)
	if(err != nil){
		log.Fatal(err)
		http.Error(res, err.Error(), 500)
		return
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(companyId)
}

func(companyController ICompanyController) GetCompaniesController(res http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	companies, err := companyController.companyDb.GetCompanies()
	if(err != nil){
		log.Fatal(err)
		http.Error(res, err.Error(), 500)
		return
	}

res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(companies)
}