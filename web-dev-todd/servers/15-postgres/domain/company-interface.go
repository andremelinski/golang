package domain

import "github.com/andremelinski/web-dev-todd/servers/15-postgres/domain/interfaces"

type ICompany interface {
	CreateCompany(companyProps interfaces.ICompanyProps) (int64, error)
}