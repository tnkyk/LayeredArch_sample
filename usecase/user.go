package usecase

import (
	"context"
	"log"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
)

type UserUsecase interface {
	UserGetAll(ctx context.Context) ([]model.User, error)
	UserGetByName(ctx context.Context, name string) (*model.User, error)
	UpsertUser(ctx context.Context, id string, name string, email string, password string) error
	DeleteUser(ctx context.Context, id string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUsecase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu *userUseCase) UserGetAll(ctx context.Context) ([]model.User, error) {
	todos, err := uu.userRepository.GetAll(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todos, nil
}

func (uu *userUseCase) UserGetByName(ctx context.Context, name string) (*model.User, error) {
	todo, err := uu.userRepository.GetByName(ctx, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, err
}

func (uu *userUseCase) UpsertUser(ctx context.Context, id string, name string, email string, password string) error {
	err := uu.userRepository.UpsertUser(ctx, id, name, email, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (uu *userUseCase) DeleteUser(ctx context.Context, id string) error {
	err := uu.userRepository.DeleteUser(ctx, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
