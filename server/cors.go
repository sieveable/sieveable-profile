package main

import (
	"net/http"
)

type CORSPolicy struct {
	next           http.Handler
	allowedOrigins []string
	allowedMethods []string
}

func NewCORSPolicy(next http.Handler, allowedOrigins []string, allowedMethods []string) *CORSPolicy {
	return &CORSPolicy{next: next, allowedOrigins: allowedOrigins,
		allowedMethods: allowedMethods}
}

func contains(element string, list []string) bool {
	for _, val := range list {
		if val == element {
			return true
		}
	}
	return false
}

// A middleware that allows requests from a single host, specified in the HTTP Host header
func (c *CORSPolicy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if contains(r.Method, c.allowedMethods) && len(r.Header["Origin"]) == 1 &&
		contains(r.Header["Origin"][0], c.allowedOrigins) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])
		c.next.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden: Disallowed origin.", 403)
	}
}
