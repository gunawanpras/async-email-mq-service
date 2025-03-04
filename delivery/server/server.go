package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gunawanpras/async-email-mq-service/config"
	"github.com/gunawanpras/async-email-mq-service/internal/setup"
	"github.com/gunawanpras/async-email-mq-service/pkg/response"
	"github.com/gunawanpras/async-email-mq-service/pkg/util/constant"
)

func Up(handler setup.Handler, config config.ServerConfig) {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return response.Error(c, err.Error(), err, constant.GenericHttpStatusMappings)
			},
		},
	)

	app.Use(cors.New())

	NewRouter(app, handler)

	err := app.Listen(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Panic(err)
	}
}
