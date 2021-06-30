package route

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/project/controller"
	"ocg.com/project/util"
)

func AuthConfig(router *fiber.Router) {
	(*router).Post("/register", controller.Register)
	(*router).Post("/login", controller.Login)
	(*router).Post("/logout", controller.Logout)
	(*router).Get("/users", controller.User)
}

func ProductsConfig(router *fiber.Router) {
	(*router).Get("/products", controller.GetAllProducts)
	(*router).Get("/products/:id", controller.GetProductById)
	(*router).Post("/products", controller.CreateNewProduct)
	(*router).Delete("/products/:id", controller.DeleteProductById)
	(*router).Get("/products/categories/:name", controller.GetProductByCategory)
	(*router).Get("/best-products", controller.GetProductsBestPrice)
	(*router).Get("/products/search/:key", controller.SearchProducts)
}

func CategoriesConfig(router *fiber.Router) {
	(*router).Get("/categories", controller.GetAllCategories)
}

func UploadConfig(router *fiber.Router) {
	(*router).Post("/upload", util.Upload)
	(*router).Static("/uploads", "./uploads")
}

func OrderConfig(router *fiber.Router) {
	(*router).Get("/orders/:id", controller.GetOrderById)
	(*router).Post("/orders", controller.CreateNewOder)
}
