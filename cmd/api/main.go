package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"web-go/configs"
	"web-go/internal/handlers"
)

func main() {
	config := configs.LoadConfig()
	logger := configs.NewLogger(os.Stderr, config.GetLevelLog())
	slog.SetDefault(logger)

	slog.Info(fmt.Sprintf("servidor rodando na porta %s", config.ServerPort))
	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("./views/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteHandler := handlers.NewNoteHandler()
	mux.Handle("/", handlers.HandlerWithError(noteHandler.NoteList))
	mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	mux.Handle("/note/new", handlers.HandlerWithError(noteHandler.NoteNew))
	mux.Handle("/note/create", handlers.HandlerWithError(noteHandler.NoteCreate))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux); err != nil {
		panic(err)
	}
}
