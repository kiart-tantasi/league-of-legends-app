package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// routing
	http.Handle("/riot/account/v1/accounts/by-riot-id/", loggingMiddleware(http.HandlerFunc(accountHandleFn)))
	http.Handle("/lol/match/v5/matches/by-puuid/", loggingMiddleware(http.HandlerFunc(matchIdsHandleFn)))
	http.Handle("/lol/match/v5/matches/", loggingMiddleware(http.HandlerFunc(matchDetailHandleFn)))
	// start
	port := "8090"
	fmt.Println("app is listening and serving on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

// handlers
func accountHandleFn(w http.ResponseWriter, r *http.Request) {
	err := sendJson(&w, "./internal/mock/accountResponse.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 500)
	}
}
func matchIdsHandleFn(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		count = 5
	}
	matchIdsMock := make([]string, count)
	for i := 0; i < count; i++ {
		matchIdsMock[i] = fmt.Sprintf("MOCK_%d", i)
	}
	bytes, err := json.Marshal(matchIdsMock)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 500)
	} else {
		w.Write(bytes)
	}
}
func matchDetailHandleFn(w http.ResponseWriter, r *http.Request) {
	err := sendJson(&w, "./internal/mock/matchDetailResponse.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), 500)
	}
}

func sendJson(w *http.ResponseWriter, filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	(*w).Header().Set("Content-Type", "application/json")
	_, err = (*w).Write(bytes)
	if err != nil {
		http.Error((*w), fmt.Sprintf("%s", err), 500)
		return err
	}
	return nil
}

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s, %d ms\n", r.Method, r.URL, time.Since(start).Milliseconds())
	}
	return http.HandlerFunc(fn)
}
