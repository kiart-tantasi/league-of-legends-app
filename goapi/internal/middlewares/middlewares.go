package middlewares

import "net/http"

// TODO: create sub-package for each middleware
func ApiMiddlewares(next http.Handler) http.Handler {
	return requestContextInitiator(userData(serverTime(userCookie(next))))
}
