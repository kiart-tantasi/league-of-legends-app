package main

import (
	"context"
	"fmt"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/internal/middlewares"
	"net/http"
	"os"

	"github.com/kiart-tantasi/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// env
	environment := os.Getenv("ENV")
	projectRoot := os.Getenv("PROJECT_ROOT")
	env.LoadEnvFile(environment, projectRoot)

	// === DEBUGGING === //
	initMongoDB(os.Getenv("MONGODB_URI"))
	// === DEBUGGING === //

	// routing
	http.Handle("/api/health", &health.HealthHandler{})
	http.Handle("/api/v1/matches", middlewares.ApiMiddlewares((http.Handler(&match.MatchHandler{}))))
	// start
	port := os.Getenv("SERVER_PORT")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

// move to separate package
// declare Client as public var
func initMongoDB(uri string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close connection
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
			return
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("lol-caching").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
