package postgre

import (
	"awesomeProject/internal/models"
	"context"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) Create(ctx context.Context, book *models.Book) (string, error) {
	result := r.DB.WithContext(ctx).Create(&book)
	return book.ID, result.Error
}

func (r *BookRepository) GetByAuthor(ctx context.Context, author string) (models.Book, error) {
	var res models.Book
	err := r.DB.WithContext(ctx).Where("author = ?", author).First(&res).Error
	if err != nil {
		return models.Book{}, err
	}
	return res, nil
}

func (r *BookRepository) GetByID(ctx context.Context, id string) (models.Book, error) {
	var res models.Book
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		return models.Book{}, err
	}
	return res, nil
}

//func (r *BookRepository) Update(ctx context.Context, ID string, book *models.Book) (models.Book, error) {
//	_ = r.DB.Save(&book)
//	return *book, nil
//}

func (r *BookRepository) Delete(ctx context.Context, author string) error {
	var res models.Book
	return r.DB.WithContext(ctx).Where("author = ?", author).Delete(&res).Error
}

func (r *BookRepository) List(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.DB.WithContext(ctx).Find(&books).Error
	return books, err
}
