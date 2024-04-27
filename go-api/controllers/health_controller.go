package controllers

import (
	"net/http"
)

type HealthController struct{}

func (hc *HealthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

