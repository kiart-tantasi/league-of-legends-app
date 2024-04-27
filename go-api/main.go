package main

import (
	"fmt"
	"net/http"
	"go-api/controllers"
)

func main() {
	hc := &(controllers.HealthController{})

	http.HandleFunc("/api/health", hc.GetHealth)

	fmt.Println("app is listening and serving")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

