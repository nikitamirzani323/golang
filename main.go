package main

import "github.com/gofiber/fiber/v2"

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome Togel")
}

func main() {
	app := fiber.New()
	app.Get("/", hello)

	app.Listen(":3000")
}
