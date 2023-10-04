package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"apiWithPostgres/modules/handlers"
	"apiWithPostgres/shared/configs"
)

func main(){
	err := configs.Load()

	if err != nil {
		panic(err)
	}
	log.Println("Api running on port 9000")

		r := chi.NewRouter()
		r.Post("/createTodo", handlers.Create)
		r.Put("/updateTodo/{id}", handlers.Update)
		r.Delete("/deleteTodo/{id}", handlers.Delete)
		r.Get("/listAll", handlers.List)
		r.Get("/getTodo/{id}", handlers.Get)

		http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()),r)
}

