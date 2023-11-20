package data

import (
	"database/sql"
	"fmt"

	"github.com/andremelinski/web-dev-todd/servers/15-postgres/domain/interfaces"
)

type DbCompanyRepo struct {
	db *sql.DB
}

func NewDbCompanyRepo(db *sql.DB)*DbCompanyRepo{
	return &DbCompanyRepo{
	db,
	}
}

func(companyRepo DbCompanyRepo) CreateCompany(companyProps interfaces.ICompanyProps)(int64, error){
	newCompany, err := companyRepo.db.Exec("INSERT INTO company(name) VALUES (?)", companyProps.Name)
	if err != nil {
		return 0, err
	}

	companyId, err := newCompany.LastInsertId()
	if err != nil {
        return 0, fmt.Errorf("LastInsertId in addAlbum: %v", err)
    }
	return companyId, nil
}

func(companyRepo DbCompanyRepo) GetCompanies()([]interfaces.ICompanyDb, error){
	companies := []interfaces.ICompanyDb{}
	rows, err := companyRepo.db.Query("SELECT * FROM company")
	defer rows.Close() 
	if err != nil {
		return companies, err
	}

	for rows.Next(){
		company := interfaces.ICompanyDb{}
		if err := rows.Scan(&company.Id, &company.Name); err != nil {
            return nil, err
        }
        companies = append(companies, company)
	}

	return companies, nil

}