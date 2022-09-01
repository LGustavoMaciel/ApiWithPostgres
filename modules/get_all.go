package models

import (
	"apiWithPostgres/modules/domain"
	"apiWithPostgres/shared/db"
)

func GetAll() (todos []domain.Todo, err error){
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}
	for rows.Next(){
		var todo domain.Todo
	err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil{
			continue
		}
		todos = append(todos, todo)
	}
	return
}