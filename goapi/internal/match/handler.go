package match

import (
	"goapi/internal/middlewares/requestcontext"
	"log"
	"net/http"
)

type MatchHandler struct{}

func (*MatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate required query parameters
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		requestcontext.WriteStatus(w, http.StatusBadRequest, r)
		return
	}
	// get matches from riot api
	matches, err := getMatchesV1(gameName, tagLine)
	if err != nil {
		log.Println("getMatchesV1 error:", err)
		requestcontext.WriteStatus(w, http.StatusBadRequest, r)
		return
	}
	// write response
	w.Header().Set("Content-Type", "application/json")
	requestcontext.WriteStatus(w, http.StatusOK, r)
	w.Write(matches)
}
