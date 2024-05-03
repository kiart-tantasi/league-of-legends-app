package middlewares

import (
	"fmt"
	"go-api/internal/match"
	"net/http"
	"time"
)

func ServerTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		statusCode := r.Context().Value(match.StatusCode)
		fmt.Printf("%d, %s, %d ms\n", statusCode, r.URL, (time.Since(start).Milliseconds()))
	}
	return http.HandlerFunc(handlerFn)
}
