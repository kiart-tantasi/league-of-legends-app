package match

import (
	"net/http"
)

type MatchHandler struct{}

func (matchHandler *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		http.Error(w, "", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain") // omittable
	matches, err := getMatches(gameName, tagLine)
	if err != nil {
		http.Error(w, "", 400)
	}
	w.Write([]byte(matches))
}
