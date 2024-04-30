package main

import (
	"fmt"
	"go-api/controllers"
	v1 "go-api/controllers/v1"
	"go-api/utils"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// env
	setUpEnv()

	// routing
	healthController := &controllers.HealthController{}
	matchController := &v1.MatchController{}
	http.HandleFunc("/api/health", healthController.GetHealth)
	http.HandleFunc("/api/matches", matchController.GetMatches)

	// start
	port := utils.GetEnv("SERVER_PORT", "8080")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

func setUpEnv() {
	env := utils.GetEnv("ENV", "development")
	if env == "production" {
		godotenv.Load(".env.production")
	}
	fmt.Printf("running with profile \"%s\"\n", env)
}
