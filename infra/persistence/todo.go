package persistence

//infra層は技術的関心事を行う層

import (
	"context"
	"log"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
	"github.com/tnkyk/LayeredArch_sample/infra/config"
)

type TodoPersistence struct {
}

func NewTodoPersistence() repository.TodoRepository {
	return &TodoPersistence{}
}

//TodoPersistence構造体がGetAllを実装しているという意
func (tp TodoPersistence) GetAll(context.Context) ([]*model.Todo, error) {
	rows, err := config.DB.Query("SELECT * FROM todos")
	if err != nil {
		//TODO: error handling
		log.Println(err)
		return nil, err
	}
	todo := &model.Todo{}
	todos := []*model.Todo{}

	for rows.Next() {
		err = rows.Scan(&todo.Id, &todo.Title, &todo.CreatedAt)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (tp TodoPersistence) GetById(ctx context.Context, id string) (*model.Todo, error) {
	row := config.DB.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	todo := &model.Todo{}
	err := row.Scan(&todo.Id, &todo.Title, &todo.CreatedAt)
	if err != nil {
		log.Println("can't get row")
		return nil, err
	}

	return todo, nil

}
