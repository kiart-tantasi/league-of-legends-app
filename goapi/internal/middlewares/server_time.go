package middlewares

import (
	"context"
	"goapi/internal/contexts"
	"log"
	"net/http"
	"time"
)

func serverTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, _r *http.Request) {
		start := time.Now()

		// init request context and store in request
		initialRequestContext := context.WithValue(_r.Context(), contexts.RequestContextKey, &contexts.RequestContext{})
		r := _r.WithContext(initialRequestContext)

		next.ServeHTTP(w, r)

		// read request context after serving
		requestContext, ok := r.Context().Value(contexts.RequestContextKey).(*contexts.RequestContext)
		if ok && requestContext != nil {
			log.Printf("%s %s, %d ms, %d\n", r.Method, r.URL, time.Since(start).Milliseconds(), requestContext.StatusCode)
		}
	}
	return http.HandlerFunc(handlerFn)
}
