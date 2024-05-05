package middlewares

import (
	"fmt"
	"go-api/internal/contexts"
	"net/http"
	"time"
)

func serverTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		statusCode := r.Context().Value(contexts.StatusCode)
		fmt.Printf("%s %s, %d ms, %d\n", r.Method, r.URL, time.Since(start).Milliseconds(), statusCode)
	}
	return http.HandlerFunc(handlerFn)
}
