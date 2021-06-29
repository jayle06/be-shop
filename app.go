package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"ocg.com/project/database"
	"ocg.com/project/route"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	baseURL := app.Group("api/v1")
	route.ProductsConfig(&baseURL)
	route.CategoriesConfig(&baseURL)
	route.AuthConfig(&baseURL)
	route.UploadConfig(&baseURL)
	route.OrderConfig(&baseURL)

	app.Listen(":3000")
}
