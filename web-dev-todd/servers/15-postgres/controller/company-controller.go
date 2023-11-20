package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/andremelinski/web-dev-todd/servers/15-postgres/repository/data"
	"github.com/julienschmidt/httprouter"
)

type ICompanyController struct {
	companyDb data.DbCompanyRepo
}

func InitCompanyController(db *sql.DB)*ICompanyController{
	return &ICompanyController{
		companyDb: *data.NewDbCompanyRepo(db),
	}
}

func(companyController ICompanyController) CreateCompanyController(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	reqBody, _ := io.ReadAll(req.Body)
	companyInfo := data.ICompanyProps{}
	
	if err := json.Unmarshal(reqBody, &companyInfo); err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	companyId, err := companyController.companyDb.CreateCompany(companyInfo)
	if(err != nil){
		log.Fatal(err)
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(companyId)
}