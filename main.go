package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/golang/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome Togel")
}

func Routers(app *fiber.App){
	app.Get("/user" user.getUsers)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)

	app.Listen(":3000")
}
