package main

import (
	"github.com/TranThaiNgoc/Fiber-GORM/database"
	"github.com/TranThaiNgoc/Fiber-GORM/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	SetupRoutes(app)
	app.Listen(":3000")
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to fiber")
}

func SetupRoutes(app *fiber.App) {
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}
