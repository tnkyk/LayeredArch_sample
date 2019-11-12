package persistence

//infra層は技術的関心事を行う層

import (
	"context"
	"fmt"
	"time"

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
func (tp *TodoPersistence) GetAll(context.Context) ([]model.Todo, error) {
	rows, err := config.DB.Query("SELECT * FROM todos")
	if err != nil {
		sqlerr := NewSQLError(err, "Can't get rows in GetAll")
		return nil, fmt.Errorf("GetAll Error >>> %w", sqlerr)
	}
	var todo model.Todo //todo->値 　&todo->address
	var todos []model.Todo

	for rows.Next() {
		err = rows.Scan(&todo.Id, &todo.Title, &todo.CreatedAt)
		if err != nil {
			sqlerr := NewSQLError(err, "Can't Scan in GetAll")
			return nil, fmt.Errorf("GetAll Error >>> %w", sqlerr)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

//nilはpointer型にconvertできない
func (tp *TodoPersistence) GetById(ctx context.Context, id string) (*model.Todo, error) {
	row := config.DB.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	var todo model.Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.CreatedAt)
	if err != nil {
		sqlerr := NewSQLError(err, "Can't Scan in GetById")
		return nil, fmt.Errorf("GetById Error >>> %w", sqlerr)
	}

	return &todo, nil

}

func (tp *TodoPersistence) UpsertTodo(ctx context.Context, id string, title string, createdAt time.Time) error {
	stmt, err := config.DB.Prepare(`INSERT INTO todos (id, title, created_at) VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE title=?, created_at=?`)
	if err != nil {
		sqlerr := NewSQLError(err, "Can't Prepare in UpsertTodo")
		return fmt.Errorf("%w", sqlerr)
	}
	_, err = stmt.Exec(id, title, createdAt, title, createdAt)
	if err != nil {
		sqlerr := NewSQLError(err, "Can't Exec in UpsertTodo")
		return fmt.Errorf("UpsertTodo Error >>> %w", sqlerr)
	}
	return nil
}

func (tp *TodoPersistence) DeleteTodo(ctx context.Context, id string) error {
	stmt, err := config.DB.Prepare(`DELETE FROM todos WHERE id = ?`)
	if err != nil {
		sqlerr := NewSQLError(err, "Can't Prepare in DeleteTodo")
		return fmt.Errorf("DeleteTodo Error >>> %w", sqlerr)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		sqlerr := NewSQLError(err, "Can't Exec in DeleteTodo")
		return fmt.Errorf("DeleteTodo Error >>> %w", sqlerr)
	}
	return nil
}
