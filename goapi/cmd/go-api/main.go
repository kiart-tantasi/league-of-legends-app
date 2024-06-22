package main

import (
	"fmt"
	"goapi/internal/cache"
	"goapi/internal/health"
	"goapi/internal/match"
	"goapi/internal/middlewares"
	"log"
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
		cache.InitMongoClient(os.Getenv("CACHE_MONGODB_URI"))
	}
	// routing
	http.Handle("/api/health", &health.HealthHandler{})
	http.Handle("/api/v1/matches", middlewares.ApiMiddlewares((http.Handler(&match.MatchHandler{}))))
	// start
	port := os.Getenv("SERVER_PORT")
	log.Println("app is listening and serving on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
