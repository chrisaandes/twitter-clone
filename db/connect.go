package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnect init */
var MongoConnect = InitConnect()

/*InitConnect ...*/
func InitConnect() *mongo.Client {
	var username = "admin"
	var pw = "kQm9i4G4TXxDIWS1"
	var host1 = "cluster0-kaoso.mongodb.net"
	var mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", username, pw, host1)
	var clientOptions = options.Client().ApplyURI(mongoURI)

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
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
