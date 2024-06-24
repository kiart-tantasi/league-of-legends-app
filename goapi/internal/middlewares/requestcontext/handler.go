package requestcontext

import (
	"context"
	"net/http"
)

func Handler(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, _r *http.Request) {
		initialRequestContext := context.WithValue(_r.Context(), RequestContextKey, &RequestContext{})
		r := _r.WithContext(initialRequestContext)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handlerFn)
}
