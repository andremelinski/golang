package data

import (
	"database/sql"
	"fmt"
)

type DbCompanyRepo struct {
	db *sql.DB
}

func NewDbCompanyRepo(db *sql.DB)*DbCompanyRepo{
	return &DbCompanyRepo{
	db,
	}
}

type ICompanyProps struct {
	Name string `json:name`
}

type ICompanyDb struct {
	Id int64 `json:id`
	ICompanyProps
}

func(companyRepo DbCompanyRepo) CreateCompany(companyProps ICompanyProps)(int64, error){
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