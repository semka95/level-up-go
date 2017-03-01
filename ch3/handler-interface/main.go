package main

import (
	"fmt"
	"net/http"
	"time"
)

// UptimeHandler writes a number of seconds since starting the response
type UptimeHandler struct {
	Started time.Time
}

// NewUptimeHandler returns a time when server starts
func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("Current uptime: %s", time.Since(h.Started)))
}

// SecretTokenHandler secures a request with a secret token
type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("secret_token") == h.secret {
		h.next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "Secret",
	})
	http.ListenAndServe(":3000", nil)
}
