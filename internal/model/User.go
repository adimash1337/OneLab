package model

type User struct {
	ID       int     `json:"id"`
	Name     *string `json:"name" validate:"required,min=2,max=30"`
	UserName *string `json:"user_name"  validate:"required,min=2,max=30"`
	Password *string `json:"password"   validate:"required,min=6"`
}

type UserRepository interface {
	Create(user User) error
	Find(UserName string) (User, error)
	Delete(UserName string) error
}
