package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnect init */
var MongoConnect = InitConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:kQm9i4G4TXxDIWS1@cluster0-kaoso.mongodb.net/test?retryWrites=true&w=majority")

/*InitConnect ...*/
func InitConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("connection succesful")
	return client

}

/*CheckConnection ...*/
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return true
	}
	return false
}
