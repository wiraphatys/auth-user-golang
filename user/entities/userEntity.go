package entities

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserResponse struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
