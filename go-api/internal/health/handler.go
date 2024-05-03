package health

import (
	"net/http"
)

type HealthHandler struct{}

func (*HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
