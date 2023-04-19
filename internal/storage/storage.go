package storage

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/models"
	"awesomeProject/internal/storage/postgre"
	"context"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByUsername(ctx context.Context, username string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Delete(ctx context.Context, username string) error
	UpdatePassword(ctx context.Context, username string, password string) error
}

type IBookRepository interface {
	GetByAuthor(ctx context.Context, author string) (models.Book, error)
	GetByID(ctx context.Context, id string) (models.Book, error)
	Create(ctx context.Context, book *models.Book) (string, error)
	Delete(ctx context.Context, author string) error
	List(ctx context.Context) ([]models.Book, error)
}
type INoteRepository interface {
	Create(ctx context.Context, note *models.Note) (string, error)
	Get(ctx context.Context, ID string) (models.Note, error)
	Delete(ctx context.Context, ID string) error
	List(ctx context.Context) ([]models.Note, error)
}

type Storage struct {
	Pg   *gorm.DB
	User IUserRepository
	Book IBookRepository
	Note INoteRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {

	pgDb, err := postgre.Dial(*cfg)
	if err != nil {
		logger.Logger().Println(err)
		return nil, err
	}

	uRepo := postgre.NewUserRepository(pgDb)
	bRepo := postgre.NewBookRepository(pgDb)
	nRepo := postgre.NewNotesRepository(pgDb)

	storage := Storage{
		User: uRepo,
		Book: bRepo,
		Note: nRepo,
	}
	return &storage, nil
}
