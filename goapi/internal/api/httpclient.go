package api

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(5000 * time.Millisecond),
	}
}
