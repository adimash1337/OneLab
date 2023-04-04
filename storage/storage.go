package storage

import (
	"awesomeProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)


// Почему не вынес в отдельную папку inMemory ? 
// Почему нету разделения на сущность / файл manager для инициализации репы
// 
type InMemoryUserRepository struct {
	users map[int]*model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*model.User),
	}
}

type UserRepository interface {
	GetByID(id int) (*model.User, error) // id => ID 
	Create(user *model.User) error
}

func (repo *InMemoryUserRepository) GetByID(id int) gin.HandlerFunc {
	return func(c *gin.Context) { // Откуда тут gin.Context ? почему у тебя хендлер на уровне репозитория ? 
		var user model.User
		var founduser model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// почему это здесь а не в отедльной функции ? 
		// -------- Приведи пожаулйста в порядок в соответствии с чистой архитектурой. Дальше не смотрел. 
		PasswordIsValid, msg := func(userpassword string, givenpassword string) (bool, string) {
			err := bcrypt.CompareHashAndPassword([]byte(givenpassword), []byte(userpassword))
			valid := true
			msg := ""
			if err != nil {
				msg = "Login Or Passowrd is Incorerct"
				valid = false
			}
			return valid, msg
		}(*user.Password, *founduser.Password)
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		if user, ok := repo.users[id]; ok {
			c.JSON(http.StatusFound, user)
		}
	}
}

func (repo *InMemoryUserRepository) Create(u *model.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var Validate = validator.New()
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		password := func(password string) string {
			bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
			if err != nil {
				log.Panic(err)
			}
			return string(bytes)
		}(*user.Password)

		user.Password = &password

		if _, ok := repo.users[u.ID]; ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with id %d already exists"})
		}
		repo.users[u.ID] = u
		c.JSON(http.StatusCreated, "Successfully created!!")
	}
}
