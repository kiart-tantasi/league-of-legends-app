package usercookie

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Handler(next http.Handler) http.Handler {
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
		message := fmt.Sprintf("user id: %s, %s", userId, getHeaderValues(r))
		log.Println(message)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handlerFn)
}

func getHeaderValues(r *http.Request) string {
	s := []string{}
	for k, v := range r.Header {
		if len(v) > 0 && v[0] != "" {
			s = append(s, fmt.Sprintf("%s:%s", k, v[0]))
		}
	}
	return strings.Join(s, ";")
}
