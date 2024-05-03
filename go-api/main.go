package main

import (
	"fmt"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/pkg/env"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// env
	setUpEnv()

	// routing
	healthHandler := &health.HealthHandler{}
	matchHandler := &match.MatchHandler{}
	http.HandleFunc("/api/health", healthHandler.GetHealth)
	http.HandleFunc("/api/v1/matches", matchHandler.GetMatchesV1)

	// start
	port := env.GetEnv("SERVER_PORT", "8080")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

// TODO: update README.md about go's env vars
func setUpEnv() {
	env := env.GetEnv("ENV", "development")
	envPath := ".env"
	if env == "production" {
		projectRoot := os.Getenv("PROJECT_ROOT")
		envPath = filepath.Join(projectRoot, ".env.production")
	}
	err := godotenv.Load(envPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("running with profile \"%s\"\n", env)
}
