package service

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage"
	"awesomeProject/internal/transport/http/middleware"
	"context"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user models.User) (string, error) {
	user.ID = uuid.NewString()
	password, err := s.HashPassword(user.Password)
	if err != nil {
		return "", nil
	}
	token, refreshtoken, _ := middleware.TokenGenerator(user.Email, user.Name, user.UserName, user.ID)
	user.Token = &token
	user.Refresh_Token = &refreshtoken
	user.Password = password
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	return s.repo.User.Create(ctx, user)
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (models.User, error) {
	return s.repo.User.GetByUsername(ctx, username)
}

func (s *UserService) GetByID(ctx context.Context, id string) (models.User, error) {
	return s.repo.User.GetByID(ctx, id)
}

func (s *UserService) Login(ctx context.Context, user *models.UserAuth) (*models.UserLogin, error) {
	userFromDB, userErr := s.repo.User.GetByUsername(ctx, user.Username)
	if userErr != nil {
		return nil, userErr
	}

	checkErr := s.CheckPassword(userFromDB.Password, user.Password)
	if checkErr != nil {
		return nil, checkErr
	}

	return &models.UserLogin{
		ID:   userFromDB.ID,
		Name: userFromDB.Name,
	}, nil
}

func (s *UserService) Delete(ctx context.Context, username string) error {
	return s.repo.User.Delete(ctx, username)
}

func (s *UserService) UpdatePassword(ctx context.Context, req *models.UpdatePasswordReq, username string) error {
	userFromDB, err := s.repo.User.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	if err := s.CheckPassword(userFromDB.Password, req.CurrentPassword); err != nil {
		return err
	}

	hash, err := s.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.repo.User.UpdatePassword(ctx, userFromDB.UserName, hash)
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Logger().Println(err)
		return "", err
	}
	return string(bytes), nil
}

func (s *UserService) CheckPassword(hashedPwd, inputPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
}
