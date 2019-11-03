package repository

import (
	"context"
	"time"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
)

type DBRepository interface {
}

type TodoRepository interface {
	GetAll(context.Context) ([]model.Todo, error)
	GetById(context.Context, string) (*model.Todo, error)
	UpsertTodo(ctx context.Context, id string, title string, createdAt time.Time) error
	DeleteTodo(ctx context.Context, id string) error
}
