package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBaseConnection struct{}

func InitDataBaseConnection()*DataBaseConnection{
	return &DataBaseConnection{}
}

func (dbConn DataBaseConnection) ConnectMongoDB()(*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/node-api?authSource=admin"))
	if err != nil { return nil, cancel, err }

	database := client.Database("Album")
	return database, cancel, nil
}

// func (dbConn DataBaseConnection) ConnectMySQLDB() (*sql.DB, error) {
// 	cfg := mysql.Config{
// 		User:                 "root",
// 		Passwd:               "mypassword",
// 		Net:                  "tcp",
// 		Addr:                 "127.0.0.1:3306",
// 		DBName:               "testdb",
// 		AllowNativePasswords: true,
// 	}

// 	db, err := sql.Open("mysql", cfg.FormatDSN())

// 	db.SetConnMaxLifetime(time.Minute * 1)
// 	if err != nil {
//         return nil, err
//     }

//     err = db.Ping()
//     if err != nil {
//         return nil, err
//     }
//     fmt.Println("Connected!")
// 	return db, nil
// }