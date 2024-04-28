package main

import (
	"fmt"
	"go-api/controllers"
	v1 "go-api/controllers/v1"
	"net/http"
)

func main() {
	healthController := &controllers.HealthController{}
	matchController := &v1.MatchController{}
	http.HandleFunc("/api/health", healthController.GetHealth)
	http.HandleFunc("/api/matches", matchController.GetMatches)
	fmt.Println("app is listening and serving")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
