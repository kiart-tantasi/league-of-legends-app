package main

import (
	"fmt"
	"go-api/internal/health"
	"go-api/internal/match"
	"go-api/pkg/env"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// env
	setUpEnv()

	// routing
	http.Handle("/api/health", &health.HealthHandler{})
	http.Handle("/api/v1/matches", serverTimeMiddleware(http.Handler(&match.MatchHandler{})))

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

// middlwares
func serverTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		// TODO: implement context to store status code
		statusMock := 0
		fmt.Printf("%d, %s, %d ms\n", statusMock, r.URL, (time.Since(start).Milliseconds()))
	})
}
