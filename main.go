package main

import (
	"github.com/RashadArbab/goServer/Database"
	"github.com/RashadArbab/goServer/Routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	Database.Init()

	app := fiber.New()

	app.Static("/", "./AA-Frontend/build")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello there")
	})

	Routes.SetupRoutes(app)

	app.Listen(":5000")
}
