package main

import (
	"fmt"
	"go-api/controllers"
	v1 "go-api/controllers/v1"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// env
	env := os.Getenv("ENV")
	if env == "production" {
		godotenv.Load(".env.production")
	} else {
		env = "development"
		godotenv.Load(".env")
	}
	fmt.Printf("running with profile \"%s\"\n", env)

	// routing
	healthController := &controllers.HealthController{}
	matchController := &v1.MatchController{}
	http.HandleFunc("/api/health", healthController.GetHealth)
	http.HandleFunc("/api/matches", matchController.GetMatches)

	// start
	port := os.Getenv("SERVER_PORT")
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
