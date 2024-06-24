package userdata

import (
	"net/http"
)

func Handler(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		if isEnabled() {
			// TODO: store user data in db
		}
	}
	return http.HandlerFunc(handlerFn)
}
