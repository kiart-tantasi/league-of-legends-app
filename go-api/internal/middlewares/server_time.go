package middlewares

import (
	"context"
	"fmt"
	"go-api/internal/contexts"
	"net/http"
	"time"
)

func serverTime(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, _r *http.Request) {
		start := time.Now()

		// init request context and store in request
		ctx := context.WithValue(_r.Context(), contexts.RequestContextKey, &contexts.RequestContext{})
		r := _r.WithContext(ctx)

		next.ServeHTTP(w, r)

		// read request context after serving
		var requestContext *contexts.RequestContext = r.Context().Value(contexts.RequestContextKey).(*contexts.RequestContext)
		if requestContext != nil && requestContext.StatusCode != nil {
			fmt.Printf("%s %s, %d ms, %d\n", r.Method, r.URL, time.Since(start).Milliseconds(), *requestContext.StatusCode)
		}
	}
	return http.HandlerFunc(handlerFn)
}
