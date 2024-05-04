package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func user(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		userId := ""
		userIdCookieName := "user_id"
		existingCookie, err := r.Cookie(userIdCookieName)
		if err != nil || existingCookie == nil {
			userId = uuid.New().String()
			ninetyDays := 90 * 24 * time.Hour
			cookie := http.Cookie{
				Name:     userIdCookieName,
				Value:    userId,
				Expires:  time.Now().Add(ninetyDays),
				Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
		} else {
			userId = existingCookie.Value
		}
		fmt.Println("===================================================")
		fmt.Println("user id: " + userId)
		logHeaders(r)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handlerFn)
}

func logHeaders(r *http.Request) {
	for k, v := range r.Header {
		if v[0] != "" {
			fmt.Printf("%s: %s\n", k, v[0])
		}
	}
}
