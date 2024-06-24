package usercookie

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const userIdCookieName = "user_id"

func Handler(next http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		userId := ""
		cookie, err := r.Cookie(userIdCookieName)
		if err != nil || cookie == nil {
			newCookie := createCookie()
			http.SetCookie(w, newCookie)
			userId = newCookie.Value
		} else {
			userId = cookie.Value
		}
		message := fmt.Sprintf("user id: %s, %s", userId, getHeaderValues(r))
		log.Println(message)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handlerFn)
}

func createCookie() *http.Cookie {
	userId := uuid.New().String()
	ninetyDays := 90 * 24 * time.Hour
	return &http.Cookie{
		Name:     userIdCookieName,
		Value:    userId,
		Expires:  time.Now().Add(ninetyDays),
		Secure:   true,
		HttpOnly: true,
	}
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
