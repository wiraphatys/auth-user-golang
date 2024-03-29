package usecases

import (
	"banky/config"
	"banky/user/entities"
	"banky/user/repositories"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecaseImpl(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepository: userRepository,
	}
}

func (u *userUsecaseImpl) RegisterUser(user *entities.User) (*entities.User, error) {

	if !isEmailValid(user.Email) {
		return nil, fmt.Errorf("invalid email address")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	result, err := u.userRepository.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUsecaseImpl) SignInUser(user *entities.UserSignIn) (string, error) {
	existUser, err := u.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	// create jwt token
	cfg := config.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  existUser.ID,
		"exp": time.Now().Add(time.Hour * time.Duration(cfg.Jwt.Expiration)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
