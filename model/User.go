package model


// за валидацию молодец
// За snake_case нет ( 
type User struct {
	ID         int     `json:"id"`
	First_Name *string `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name  *string `json:"last_name"  validate:"required,min=2,max=30"`
	Password   *string `json:"password"   validate:"required,min=6"`
}
