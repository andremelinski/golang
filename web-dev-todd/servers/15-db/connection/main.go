package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
 // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "mypassword",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "testdb",
		AllowNativePasswords: true,
    }

	db, err = sql.Open("mysql", cfg.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 1)
	if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}