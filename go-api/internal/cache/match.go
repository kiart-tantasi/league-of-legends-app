package cache

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CacheMatchDetail(id, responseBody string) error {
	collection := client.Database(databaseName).Collection("matches")
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"responseBody", responseBody}}}}
	_, err := collection.UpdateByID(context.TODO(), id, update, opts)
	return err
}
