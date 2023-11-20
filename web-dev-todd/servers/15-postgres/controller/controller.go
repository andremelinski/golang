package controller

import "database/sql"

type Controllers struct {
	CompanyController *ICompanyController
	TablesDb *CreateDbController
}

func InitControllers(db *sql.DB) *Controllers {
	createTables := NewCreateTables(db)
	companyController := InitCompanyController(db)
	return &Controllers{
		CompanyController: companyController,
		TablesDb: createTables,
	}
}