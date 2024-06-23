package middlewares

import (
	"goapi/internal/contexts"
	"log"
	"net/http"
	"time"
)

func serverTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		requestContext, ok := r.Context().Value(contexts.RequestContextKey).(*contexts.RequestContext)
		if ok && requestContext != nil {
			log.Printf("%s %s, %d ms, %d\n", r.Method, r.URL, time.Since(start).Milliseconds(), requestContext.StatusCode)
		}
	}
	return http.HandlerFunc(handlerFn)
}
