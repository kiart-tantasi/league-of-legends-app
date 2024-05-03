package match

import (
	"fmt"
	"net/http"
)

type MatchHandler struct{}

func (matchHandler *MatchHandler) GetMatchesV1(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		http.Error(w, "", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain") // omittable
	matches, err := getMatchesV1(gameName, tagLine)
	if err != nil {
		fmt.Println("GetMatches error:", err)
		http.Error(w, "", 400)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(matches)
}
