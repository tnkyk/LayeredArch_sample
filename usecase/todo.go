package usecase

//usecaseは業務の手順を簡潔に記す場所

import (
	"context"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
)

type TodoUseCase interface {
	GetAll(context.Context) ([]*model.Todo, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository
}

//ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewTodoUseCase(tr repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

//Todoデータを全件取得するためのユースケース
func (tu todoUseCase) GetAll(ctx context.Context) (todos []*model.Todo, err error) {
	// Persistenceを呼出
	todos, err = tu.todoRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
