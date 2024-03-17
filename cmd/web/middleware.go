package main

import (
	"net/http"
)

func securityHeadersMiddleware(next http.Handler) http.Handler {
	// note: middleware has to return a handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Preprocessing----------------------------------------------------
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		//call the original handler after adding the two headers //continue the chain
		next.ServeHTTP(w, r)
	})
}
