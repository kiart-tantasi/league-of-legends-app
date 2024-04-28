package controllers

import (
	"go-api/services"
	"net/http"
)

type MatchController struct{}

func (matchController *MatchController) GetMatches(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	if gameName == "" || tagLine == "" {
		http.Error(w, "", 400)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte((&services.MatchService{}).GetMatches("GAME_NAME_MOCK", "TAG_LINE_MOCK")))
	w.WriteHeader(http.StatusBadRequest)
}
