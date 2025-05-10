package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jbeausoleil/go-webapp/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	taskHandler := handlers.NewTaskController()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.List)
		r.Get("/{taskID}", taskHandler.Show)
		r.Post("/", taskHandler.Create)
	})

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
