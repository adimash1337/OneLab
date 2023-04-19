package postgre

import (
	"awesomeProject/internal/models"
	"context"
	"gorm.io/gorm"
)

type INotesRepository interface {
	Create(ctx context.Context, note *models.Note) (string, error)
	Get(ctx context.Context, ID string) (models.Note, error)
	Delete(ctx context.Context, ID string) error
	List(ctx context.Context) ([]models.Note, error)
}

type NotesRepository struct {
	DB *gorm.DB
}

func NewNotesRepository(DB *gorm.DB) *NotesRepository {
	return &NotesRepository{DB: DB}
}

func (r *NotesRepository) Create(ctx context.Context, note *models.Note) (string, error) {
	result := r.DB.WithContext(ctx).Omit("deleted_at", "date_returned").Create(&note)
	return note.ID, result.Error
}

func (r *NotesRepository) Get(ctx context.Context, ID string) (models.Note, error) {
	var note models.Note
	err := r.DB.WithContext(ctx).Where("id = ?", ID).First(&note).Error
	if err != nil {
		return models.Note{}, err
	}
	return note, nil
}

func (r *NotesRepository) Delete(ctx context.Context, ID string) error {
	var note models.Note
	return r.DB.WithContext(ctx).Where("id = ?", ID).Delete(&note).Error
}

func (r *NotesRepository) List(ctx context.Context) ([]models.Note, error) {
	var records []models.Note
	err := r.DB.WithContext(ctx).Find(&records)
	return records, err.Error
}
