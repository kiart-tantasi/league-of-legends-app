package match

import (
	"fmt"
	"go-api/internal/contexts"
	"net/http"
)

type MatchHandler struct{}

func (*MatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		contexts.WriteHeaderAndContext(w, http.StatusBadRequest, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain") // omittable
	matches, err := getMatchesV1(gameName, tagLine)
	if err != nil {
		fmt.Println("GetMatches error:", err)
		contexts.WriteHeaderAndContext(w, http.StatusBadRequest, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	contexts.WriteHeaderAndContext(w, http.StatusOK, r)
	w.Write(matches)
}
