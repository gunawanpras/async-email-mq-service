package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateEmailTask(c *fiber.Ctx) error
}
