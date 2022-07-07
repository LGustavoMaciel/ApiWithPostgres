package main

import (
	"apiWithPostgres/configs"
	"apiWithPostgres/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main(){
	err := configs.Load()

	if err != nil {
		panic(err)
	}
	log.Println("Api running on port 9000")

		r := chi.NewRouter()
		r.Post("/", handlers.Create)
		r.Put("/{id}", handlers.Update)
		r.Delete("/{id}", handlers.Delete)
		r.Get("/", handlers.List)
		r.Get("/{id}", handlers.Get)

		http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()),r)
}

