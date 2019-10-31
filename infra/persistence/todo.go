package persistence

//infra層は技術的関心事を行う層

import (
	"context"
	"time"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
)

type TodoPersistence struct {
}

func NewTodoPersistence() repository.TodoRepository {
	return &TodoPersistence{}
}

//TodoPersistence構造体がGetAllを実装しているという意
func (tp TodoPersistence) GetAll(context.Context) ([]*model.Todo, error) {
	todo1 := model.Todo{Id: 1, Title: "a", Author: "yuki", CreatedAt: time.Now().Add(-24 * time.Hour)}
	todo2 := model.Todo{Id: 2, Title: "b", Author: "tanaka", CreatedAt: time.Now().Add(-24 * time.Hour)}

	return []*model.Todo{&todo1, &todo2}, nil
}
