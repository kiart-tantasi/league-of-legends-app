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

func setUpEnv() {
	env := env.GetEnv("ENV", "development")
	if env == "production" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		// TODO: delete after test
		fmt.Println("dir:", dir)
		godotenv.Load(filepath.Join(dir, ".env.production"))
	}
	fmt.Printf("running with profile \"%s\"\n", env)
}
