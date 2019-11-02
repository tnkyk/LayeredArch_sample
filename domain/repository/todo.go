package repository

import (
	"context"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
)

type DBRepository interface {
}

type TodoRepository interface {
	GetAll(context.Context) ([]*model.Todo, error)
	GetById(context.Context, string) (*model.Todo, error)
}
