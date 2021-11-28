package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

// WriteToConsole Logger is a middleware that logs the request as it goes in and the response as it goes out.
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// CSRFManager is a middleware that provides protection against Cross-Site Request Forgery (CSRF) attacks.
func CSRFManager(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.Debug,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads the session from the request.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
