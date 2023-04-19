package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            string         `json:"id" gorm:"serial;primary_key"`
	Name          string         `json:"name" validate:"required,min=2,max=30"`
	UserName      string         `json:"username" gorm:"unique" validate:"required,min=2,max=30"`
	Password      string         `json:"password"   validate:"required,min=6"`
	Email         string         `json:"email" gorm:"unique"`
	Token         *string        `json:"token"`
	Refresh_Token *string        `josn:"refresh_token"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UpdatedAt     time.Time      `gorm:"autoCreateTime"`
	CreatedAt     time.Time
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdatePasswordReq struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
