package middlewares

import "net/http"

func ApiMiddlewares(next http.Handler) http.Handler {
	return serverTime(user(next))
}
