package persistence

import (
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
)

type TodoPersistence struct {
}

func NewTodoPersistence() repository.TodoRepository {
	return &TodoPersistence{}
}
