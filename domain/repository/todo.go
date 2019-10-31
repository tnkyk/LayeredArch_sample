package repository

import (
	"context"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
)

type TodoRepository interface {
	GetAll(context.Context) ([]*model.Todo, error)
}
