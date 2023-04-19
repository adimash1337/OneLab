package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"context"
	"github.com/google/uuid"
	"time"
)

type NoteService struct {
	repo *storage.Storage
}

func NewNoteService(repo *storage.Storage) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) Create(ctx context.Context, record *models.Note, uid string, bid string) (string, error) {
	userFromDB, err := s.repo.User.GetByID(ctx, uid)
	bookFromDB, err := s.repo.Book.GetByID(ctx, bid)
	if err != nil {
		return "", err
	}
	record.ID = uuid.NewString()
	record.UserID = userFromDB.ID
	record.BookID = bookFromDB.ID
	record.DateBorrowed = time.Now()
	record.DueDate = record.DateBorrowed.AddDate(0, 0, 10)

	return s.repo.Note.Create(ctx, record)
}

func (s *NoteService) Delete(ctx context.Context, ID string) error {

	return s.repo.Note.Delete(ctx, ID)
}

func (s *NoteService) Get(ctx context.Context, ID string) (models.Note, error) {
	return s.repo.Note.Get(ctx, ID)
}

func (s *NoteService) List(ctx context.Context) ([]models.Note, error) {
	return s.repo.Note.List(ctx)
}
