package middlewares

import (
	"fmt"
	"go-api/internal/contexts"
	"net/http"
	"time"
)

func ServerTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		statusCode := r.Context().Value(contexts.StatusCode)
		fmt.Printf("%d, %s, %d ms\n", statusCode, r.URL, (time.Since(start).Milliseconds()))
	}
	return http.HandlerFunc(handlerFn)
}
