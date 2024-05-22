package main

import (
	"fmt"
	"go-api/internal/cache"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/internal/middlewares"
	"net/http"
	"os"

	"github.com/kiart-tantasi/env"
)

func main() {
	// env
	environment := os.Getenv("ENV")
	projectRoot := os.Getenv("PROJECT_ROOT")
	env.LoadEnvFile(environment, projectRoot)
	// cache
	if cache.IsEnabled() {
		cache.InitMongoClient(os.Getenv("MONGODB_URI"))
	}
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
