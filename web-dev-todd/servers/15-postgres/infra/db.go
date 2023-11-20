package infra

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DataBaseConnection struct{}

func InitDataBaseConnection()*DataBaseConnection{
	return &DataBaseConnection{}
}

func (dbConn DataBaseConnection) ConnectMySQLDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mypassword",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "testdb",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	db.SetConnMaxLifetime(time.Minute * 1)
	if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }
	return db, nil
}