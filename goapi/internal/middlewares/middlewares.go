package middlewares

import (
	"goapi/internal/middlewares/requestcontext"
	"goapi/internal/middlewares/servertime"
	"goapi/internal/middlewares/usercookie"
	"goapi/internal/middlewares/userdata"
	"net/http"
)

func ApiMiddlewares(next http.Handler) http.Handler {
	return requestcontext.Handler((userdata.Handler(servertime.Handler(usercookie.Handler(next)))))
}
