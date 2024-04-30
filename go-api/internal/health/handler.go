package health

import (
	"net/http"
)

type HealthHandler struct{}

func (hc *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
