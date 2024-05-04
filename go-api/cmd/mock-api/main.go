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
	http.Handle("/riot/account/v1/accounts/by-riot-id/", middleware(http.HandlerFunc(accountHandleFn)))
	http.Handle("/lol/match/v5/matches/by-puuid/", middleware(http.HandlerFunc(matchIdsHandleFn)))
	http.Handle("/lol/match/v5/matches/", middleware(http.HandlerFunc(matchDetailHandleFn)))
	// start
	port := "8090"
	fmt.Println("app is listening and serving on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

// handlers
func accountHandleFn(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("../../internal/mock/accountResponse.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/json")
	http.ServeContent(w, r, "", time.Now(), file)
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
		w.WriteHeader(500)
	} else {
		w.Write(bytes)
	}
}
func matchDetailHandleFn(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("../../internal/mock/matchDetailResponse.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/json")
	http.ServeContent(w, r, "", time.Now(), file)
}

// middleware
func middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("%s, %dms\n", r.URL, time.Since(start).Milliseconds())
	}
	return http.HandlerFunc(fn)
}
