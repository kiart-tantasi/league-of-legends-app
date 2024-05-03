package main

import (
	"fmt"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/internal/middlewares"
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
	http.Handle("/api/health", &health.HealthHandler{})
	http.Handle("/api/v1/matches", middlewares.ApiMiddlewares((http.Handler(&match.MatchHandler{}))))

	// start
	port := env.GetEnv("SERVER_PORT", "8080")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

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
