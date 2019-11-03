package repository

import (
	"context"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	UpsertUser(ctx context.Context, id string, name string, email string, password string) error
	DeleteUser(ctx context.Context, id string) error
}
