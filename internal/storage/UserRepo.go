package storage

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/model"
	"fmt"
)

type UserRepo struct {
	db map[string]model.User
}

func NewRepo() *UserRepo {
	return &UserRepo{
		db: make(map[string]model.User),
	}
}

func (u UserRepo) Find(UserName string) (model.User, error) {
	user, ok := u.db[UserName]
	if !ok {
		logger.Logger().Println("user with username %s not found", UserName)
	}
	return user, nil
}

func (u UserRepo) Create(user model.User) error {
	if _, ok := u.db[*user.UserName]; ok {
		return fmt.Errorf("user with username %s already exists", user.UserName)
	}
	u.db[*user.UserName] = user
	logger.Logger().Println("user created:", user)
	return nil
}

func (u UserRepo) Delete(UserName string) error {
	if _, ok := u.db[UserName]; !ok {
		logger.Logger().Println("user with username %s not found", UserName)
		return fmt.Errorf("user with username %s not found", UserName)
	}
	delete(u.db, UserName)
	return nil
}
