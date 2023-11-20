package controller

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CreateDbController struct {
	Db *sql.DB
}

func NewCreateTables(db *sql.DB)*CreateDbController{
	return &CreateDbController{
		db,
	}
}

func (dbStructure CreateDbController)CreateDbTables(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	dbStructure.Db.Exec("CREATE TABLE IF NOT EXISTS company (ID INT PRIMARY KEY NOT NULL, NAME TEXT NOT NULL, CREATED_AT timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, UPDATED_AT timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP);")

	// dbStructure.Db.Exec(`CREATE TABLE IF NOT EXISTS employees (
	// 	ID INT PRIMARY KEY NOT NULL, 
	// 	NAME TEXT NOT NULL, 
	// 	RANK INT NOT NULL, 
	// 	ADDRESS CHAR(50), 
	// 	SALARY REAL DEFAULT 25500.00, 
	// 	BDAY DATE DEFAULT '1995-12-29', 
	// 	CREATED_AT timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, 
	// 	UPDATED_AT timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP);
	// `)
}