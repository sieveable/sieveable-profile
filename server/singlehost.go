package main

import (
	"net/http"
)

type SingleHost struct {
	next        http.Handler
	allowedHost string
}

func NewSingleHost(next http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{next: next, allowedHost: allowedHost}
}

// A middleware that allows requests from a single host, specified in the HTTP Host header
func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Host == s.allowedHost {
		s.next.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}
