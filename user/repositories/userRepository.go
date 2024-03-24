package repositories

import "banky/user/entities"

type UserRepository interface {
	InsertUser(user *entities.User) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
}
