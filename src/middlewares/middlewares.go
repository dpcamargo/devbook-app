package middlewares

import (
	"app/src/cookies"
	"log"
	"net/http"
)

func Logger(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFunc(w, r)
	}
}

func Autenticar(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Ler(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		proximaFunc(w, r)
	}
}
