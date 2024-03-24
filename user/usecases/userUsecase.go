package usecases

import "banky/user/entities"

type UserUsecase interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	SignInUser(user *entities.UserSignIn) (string, error)
}
