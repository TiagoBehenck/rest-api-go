package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TiagoBehenck/rest-api-go/configs"
	"github.com/TiagoBehenck/rest-api-go/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	log.Printf("%v", configs.GetServerPort())

	// rotas
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
