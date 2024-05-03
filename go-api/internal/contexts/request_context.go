package contexts

import (
	"context"
	"net/http"
)

// https://stackoverflow.com/a/74972993/21331113
// https://stackoverflow.com/a/71686039/21331113
func WriteHeaderAndContext(w http.ResponseWriter, statusCode int, r *http.Request) {
	*r = *(r.WithContext(context.WithValue(r.Context(), StatusCode, statusCode)))
	w.WriteHeader(statusCode)
}
