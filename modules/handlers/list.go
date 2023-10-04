package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	repositories "apiWithPostgres/modules/repositories"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := repositories.GetAll()
	if err != nil {
		log.Println("Erro ao obter regitros", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
