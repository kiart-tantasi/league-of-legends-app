package contexts

import (
	"net/http"
)

type RequestContext struct {
	StatusCode *int
}

// https://stackoverflow.com/a/74972993/21331113
// https://stackoverflow.com/a/71686039/21331113
func WriteStatus(w http.ResponseWriter, statusCode int, r *http.Request) {
	w.WriteHeader(statusCode)
	// retrieve request context and write status code
	var requestContext *RequestContext = r.Context().Value(RequestContextKey).(*RequestContext)
	if requestContext != nil {
		requestContext.StatusCode = &statusCode
	}
}
