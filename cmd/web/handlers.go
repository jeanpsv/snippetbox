package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeanpsv/snippetbox/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go")

		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			"./ui/html/pages/home.tmpl.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.Logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func SnippetView(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	}
}

func SnippetCreate(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Display a form for creating a new snippet..."))
	}
}

func SnippetCreatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Save a new snippet..."))
	}
}
