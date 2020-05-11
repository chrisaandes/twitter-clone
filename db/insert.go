package db

import (
	"context"
	"time"

	"github.com/chrisaandes/twitter-clone/models"
	"github.com/chrisaandes/twitter-clone/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Insert ...*/
func Insert(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoConnect.Database("twitter")
	collection := db.Collection("users")

	u.Password, _ = utils.PasswordGenerator(u.Password)

	result, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)
	return id.String(), true, nil

}
