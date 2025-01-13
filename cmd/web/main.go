package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jeanpsv/snippetbox/config"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	config := &config.Application{
		Logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", Home(config))
	mux.Handle("GET /snippet/view/{id}", SnippetView(config))
	mux.Handle("GET /snippet/create", SnippetCreate(config))
	mux.Handle("POST /snippet/create", SnippetCreatePost(config))

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
