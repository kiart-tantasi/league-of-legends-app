package match

import (
	"context"
	"fmt"
	"net/http"
)

type MatchHandler struct{}

func (*MatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		WriteHeaderAndContext(w, 400, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain") // omittable
	matches, err := getMatchesV1(gameName, tagLine)
	if err != nil {
		fmt.Println("GetMatches error:", err)
		WriteHeaderAndContext(w, 400, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	WriteHeaderAndContext(w, 200, r)
	w.Write(matches)
}

// status code context - https://stackoverflow.com/a/74972993/21331113
type AppContext string

var StatusCode = AppContext("statusCode")

func WriteHeaderAndContext(w http.ResponseWriter, statusCode int, r *http.Request) {
	ctx := context.WithValue(r.Context(), StatusCode, statusCode)
	*r = *(r.WithContext(ctx))
	w.WriteHeader(statusCode)
}
