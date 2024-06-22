package cache

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type matchCache struct {
	Id           string `bson:"_id"`
	ResponseBody string `bson:"responseBody"`
}

func CacheMatchDetail(matchId, responseBody string) error {
	collection := getCollection()
	opts := options.Update().SetUpsert(true)
	update := bson.D{{"$set", bson.D{{"responseBody", responseBody}}}}
	ctx, cancel := context.WithTimeout(context.Background(), getDefaultTimeout())
	defer cancel()
	_, err := collection.UpdateByID(ctx, matchId, update, opts)
	return err
}

func GetMatchDetail(matchId string) (string, error) {
	collection := getCollection()
	filter := bson.D{{"_id", matchId}}
	ctx, cancel := context.WithTimeout(context.Background(), getDefaultTimeout())
	defer cancel()
	result := collection.FindOne(ctx, filter)
	var match matchCache
	err := result.Decode(&match)
	if err != nil {
		return "", err
	}
	return match.ResponseBody, nil
}

func getCollection() *mongo.Collection {
	return client.Database(databaseName).Collection("matches")
}

func getDefaultTimeout() time.Duration {
	return 1000 * time.Millisecond
}
