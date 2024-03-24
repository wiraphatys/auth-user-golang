package handlers

import (
	"banky/config"
	"banky/user/entities"
	"banky/user/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHttpHandler) RegisterUser(c *fiber.Ctx) error {
	reqBody := new(entities.User)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	result, err := h.userUsecase.RegisterUser(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	userResponse := entities.UserResponse{
		ID:        result.ID,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}

	response := NewResponse(true, "Register user successfully", userResponse)
	return SendResponse(c, response)
}

func (h *userHttpHandler) SignInUser(c *fiber.Ctx) error {
	reqBody := new(entities.UserSignIn)
	if err := c.BodyParser(reqBody); err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	token, err := h.userUsecase.SignInUser(reqBody)
	if err != nil {
		response := NewResponse(false, err.Error(), nil)
		return SendResponse(c, response)
	}

	cfg := config.GetConfig()

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * time.Duration(cfg.Jwt.Expiration)),
		HTTPOnly: true,
	})

	response := NewResponse(true, "Sign in successfully", nil)
	return SendResponse(c, response)
}
