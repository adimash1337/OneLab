package postgre

import (
	"awesomeProject/internal/models"
	"context"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByUsername(ctx context.Context, username string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Delete(ctx context.Context, book models.Book) error
	UpdatePassword(ctx context.Context, username string, password string) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) (string, error) {
	result := r.DB.WithContext(ctx).Create(&user)
	return user.ID, result.Error
}

func (r *UserRepository) UpdatePassword(ctx context.Context, username string, password string) error {
	return r.DB.Model(&models.User{}).Where("user_name = ?", username).Update("password", password).Error
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	var res models.User
	err := r.DB.WithContext(ctx).Where("user_name = ?", username).First(&res).Error
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (models.User, error) {
	var res models.User
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&res).Error
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (r *UserRepository) Delete(ctx context.Context, username string) error {
	var res models.User
	return r.DB.WithContext(ctx).Where("user_name = ?", username).Delete(&res).Error
}
