package main

import (
	"fmt"
	"go-api/env"
	"go-api/health"
	"go-api/match"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// env
	setUpEnv()

	// routing
	healthHandler := &health.HealthHandler{}
	matchHandler := &match.MatchHandler{}
	http.HandleFunc("/api/health", healthHandler.GetHealth)
	http.HandleFunc("/api/matches", matchHandler.GetMatches)

	// start
	port := env.GetEnv("SERVER_PORT", "8080")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

func setUpEnv() {
	env := env.GetEnv("ENV", "development")
	if env == "production" {
		godotenv.Load(".env.production")
	}
	fmt.Printf("running with profile \"%s\"\n", env)
}
