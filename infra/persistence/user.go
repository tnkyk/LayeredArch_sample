package persistence

import (
	"context"
	"log"

	"github.com/tnkyk/LayeredArch_sample/domain/model"
	"github.com/tnkyk/LayeredArch_sample/domain/repository"
	"github.com/tnkyk/LayeredArch_sample/infra/config"
)

//レシーバ
type UserPersistence struct {
}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}

func (up *UserPersistence) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user := model.User{}
	users := []model.User{}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (up *UserPersistence) GetByName(ctx context.Context, name string) (*model.User, error) {
	row := config.DB.QueryRow("SELECT * FROM users WHERE name = ?", name)

	var user model.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (up *UserPersistence) UpsertUser(ctx context.Context, id string, name string, email string, password string) error {
	stmt, err := config.DB.Prepare(`INSERT INTO users (id, name, email,password) VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE name=?, email=?, password=?`)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(id, name, email, password, name, email, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (up *UserPersistence) DeleteUser(ctx context.Context, id string) error {
	stmt, err := config.DB.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
