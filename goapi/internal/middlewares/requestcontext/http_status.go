package requestcontext

import (
	"goapi/internal/contexts"
	"net/http"
)

type RequestContext struct {
	StatusCode int
}

// https://stackoverflow.com/a/74972993/21331113
// https://stackoverflow.com/a/71686039/21331113
func WriteStatus(w http.ResponseWriter, statusCode int, r *http.Request) {
	w.WriteHeader(statusCode)
	// retrieve request context and write status code
	ctx, ok := r.Context().Value(contexts.RequestContextKey).(*RequestContext)
	if ok && ctx != nil {
		ctx.StatusCode = statusCode
	}
}
