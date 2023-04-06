package storage

import (
	"awesomeProject/internal/model"
)

type IUserRepo interface {
	Create(user model.User) error
	Find(UserName string) (model.User, error)
	Delete(UserName string) error
}

type Storage struct {
	IUserRepo
}

func NewStorage() *Storage {
	return &Storage{
		IUserRepo: NewRepo(),
	}
}
