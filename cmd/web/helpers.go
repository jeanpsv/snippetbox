package main

import (
	"net/http"
	"runtime/debug"

	"github.com/jeanpsv/snippetbox/config"
)

func serverError(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	trace := string(debug.Stack())

	app.Logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientError(app *config.Application, w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
