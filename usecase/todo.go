package usecase

//usecaseは業務の手順を簡潔に記す場所

import (
	"context"
	"log"
	"time"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
)

type TodoUseCase interface {
	TodoGetAll(context.Context) ([]*model.Todo, error)
	TodoGetById(context.Context, string) (*model.Todo, error)
	UpsertTodo(ctx context.Context, id string, title string, createdAt time.Time) error
	DeleteTodo(ctx context.Context, id string) error
}

type todoUseCase struct {
	todoRepository repository.TodoRepository //TodoRepositoryインターフェースを満たす必要がある
}

//ここでドメイン層のインターフェースとユースケース層のインターフェースをつなげている。
func NewTodoUseCase(tr repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

//Todoデータを全件取得するためのユースケース
func (tu todoUseCase) TodoGetAll(ctx context.Context) (todos []*model.Todo, err error) {
	// Persistenceを呼出
	todos, err = tu.todoRepository.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}

//TodoGetById：　IDを用いてTodoを取得する
func (tu todoUseCase) TodoGetById(ctx context.Context, id string) (todo *model.Todo, err error) {
	//Persistenceを呼び出し
	todo, err = tu.todoRepository.GetById(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, nil
}

func (tu todoUseCase) UpsertTodo(ctx context.Context, id string, title string, createdAt time.Time) error {
	err := tu.todoRepository.UpsertTodo(ctx, id, title, createdAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (tu todoUseCase) DeleteTodo(ctx context.Context, id string) error {
	err := tu.todoRepository.DeleteTodo(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
