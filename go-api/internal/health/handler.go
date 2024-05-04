package health

import (
	"net/http"
	"time"
)

type HealthHandler struct{}

func (*HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fake server time
	time.Sleep(500 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
}
