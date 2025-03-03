package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/async-email-mq-service/internal/setup"
)

func NewRouter(app *fiber.App, handler setup.Handler) {
	// app.Use(recover.New())
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })

	emailTasks := app.Group("/email-tasks")
	emailTasks.Post("/", handler.EmailTaskHandler.CreateEmailTask)
}
