package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go!")
	})

	app.Get("/v", func(c *fiber.Ctx) error {
		return c.SendString("v1")
	})

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
