package Routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", GetAll)
	api.Post("/", CreateUser)
	api.Get("/:id", GetSingle)
	api.Patch("/", updateNote)
}
