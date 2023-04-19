package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"context"
	"errors"
)

type IUserService interface {
	GetByUsername(ctx context.Context, username string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Login(ctx context.Context, user *models.UserAuth) (*models.UserLogin, error)
	Delete(ctx context.Context, username string) error
	UpdatePassword(ctx context.Context, req *models.UpdatePasswordReq, username string) error
}

type IBookService interface {
	GetByAuthor(ctx context.Context, author string) (models.Book, error)
	GetByID(ctx context.Context, id string) (models.Book, error)
	Create(ctx context.Context, book *models.Book) (string, error)
	Delete(ctx context.Context, author string) error
	List(ctx context.Context) ([]models.Book, error)
}

type INoteService interface {
	Create(ctx context.Context, record *models.Note, uid string, bid string) (string, error)
	Get(ctx context.Context, ID string) (models.Note, error)
	Delete(ctx context.Context, ID string) error
	List(ctx context.Context) ([]models.Note, error)
}

type Manager struct {
	User IUserService
	Book IBookService
	Note INoteService
}

func NewManager(storage *storage.Storage) (*Manager, error) {
	uSrv := NewUserService(storage)
	bSrv := NewBookService(storage)
	nSrv := NewNoteService(storage)
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	return &Manager{
		User: uSrv,
		Book: bSrv,
		Note: nSrv,
	}, nil
}
