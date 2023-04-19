package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"context"
	"github.com/google/uuid"
)

type BookService struct {
	repo *storage.Storage
}

func NewBookService(repo *storage.Storage) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) Create(ctx context.Context, book *models.Book) (string, error) {
	book.ID = uuid.NewString()
	return s.repo.Book.Create(ctx, book)
}

func (s *BookService) Delete(ctx context.Context, ID string) error {
	return s.repo.Book.Delete(ctx, ID)
}

func (s *BookService) GetByAuthor(ctx context.Context, username string) (models.Book, error) {
	return s.repo.Book.GetByAuthor(ctx, username)
}

func (s *BookService) GetByID(ctx context.Context, id string) (models.Book, error) {
	return s.repo.Book.GetByID(ctx, id)
}

func (s *BookService) List(ctx context.Context) ([]models.Book, error) {
	return s.repo.Book.List(ctx)
}
