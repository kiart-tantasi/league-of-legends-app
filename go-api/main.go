package main

import (
	"fmt"
	"go-api/internal/env"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/internal/middlewares"
	"net/http"
	"os"
)

func main() {
	// env
	env.LoadEnvFile()
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
