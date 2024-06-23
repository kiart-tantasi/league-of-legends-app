package middlewares

import (
	"context"
	"goapi/internal/contexts"
	"net/http"
)

func requestContextInitiator(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, _r *http.Request) {
		initialRequestContext := context.WithValue(_r.Context(), contexts.RequestContextKey, &contexts.RequestContext{})
		r := _r.WithContext(initialRequestContext)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handlerFn)
}
