package userdata

import (
	"net/http"
)

func Handler(next http.Handler) http.Handler {
	// TODO: config on/off
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		// TODO: store user data in db
	}
	return http.HandlerFunc(handlerFn)
}
