package handlers

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	RegisterUser(c *fiber.Ctx) error
	SignInUser(c *fiber.Ctx) error
}
