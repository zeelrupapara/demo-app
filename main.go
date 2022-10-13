package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		greeting := fmt.Sprintf("Hi, %s Have a Great Day", name)
		return c.SendString(greeting)
	})
	log.Fatal(app.Listen(":3000"))
}
