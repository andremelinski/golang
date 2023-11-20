package repository

import (
	"database/sql"

	"github.com/andremelinski/web-dev-todd/servers/15-postgres/repository/data"
)

type Repositories struct {
	CompanyRepo *data.DbCompanyRepo
	// EmployeeRepo *data.DbEmployeeRepo
}

func InitRepositories(db *sql.DB)*Repositories{
	companyRepo := data.NewDbCompanyRepo(db)
	// employee := data.NewEmployeeRepo(db)
	return &Repositories{
		companyRepo,
		// employee,

	}
}