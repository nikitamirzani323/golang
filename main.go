package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/golang/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome Togel")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/users/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)

	app.Listen(":3000")
}
