package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andremelinski/web-dev-todd/servers/15-db/mongo/infra"
	"go.mongodb.org/mongo-driver/bson"
)


func main() {
	mongodb, cancel, err := infra.InitDataBaseConnection().ConnectMongoDB()
	// ctx, cancel := context.WithCancel(context.Background())
	collection := mongodb.Collection("name")

  defer cancel()
	if(err != nil){
		log.Fatalln(err)
	}

		res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})

	if(err != nil){
		log.Fatalln(err)
	}
	id := res.InsertedID

	fmt.Println("Connected!")

		fmt.Println(id)

}