package servertime

import (
	"goapi/internal/middlewares/requestcontext"
	"log"
	"net/http"
	"time"
)

func Handler(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		requestContext, ok := r.Context().Value(requestcontext.RequestContextKey).(*requestcontext.RequestContext)
		if ok && requestContext != nil {
			log.Printf("%s %s, %d ms, %d\n", r.Method, r.URL, time.Since(start).Milliseconds(), requestContext.StatusCode)
		}
	}
	return http.HandlerFunc(handlerFn)
}
