//filename: cmd/web/middleware.go


//Written by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Tested by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Debbuged by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez

package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
)

func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "deny")

			next.ServeHTTP(w, r)
		})
}

//the func( w http.ResponseWrite..) is a regular function that is being passed
//Intercepts a response0

func (app *application) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//when the request comes to me
		start := time.Now()
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr,
			r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
		// when the response comes to me
		app.infoLog.Printf("Request took %v", time.Since(start))
	})
}

//helps to log the incoming requests and their corresponding responses, which can be useful for debugging, troubleshooting, and monitoring the web application.
func (app *application) RecoverPanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("connection", "close")
				trace := fmt.Sprintf("%s\n%s", err, debug.Stack()) //we make a string and then debug the stack
				app.errorLog.Output(2, trace)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) { //r contains the session information
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}
