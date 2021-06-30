package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/project/model"
	"ocg.com/project/repository"
)

func CreateNewProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(repository.CreateProduct(product))
}

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(repository.FindAllProducts())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(repository.FindProductById(int64(id)))
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	repository.DeleteProductById(int64(id))
	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}

func GetProductByCategory(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.JSON(repository.FindProductByCategory(name))
}

func GetProductsBestPrice(c *fiber.Ctx) error {
	return c.JSON(repository.FindTop4MinPriceProduct())
}

func SearchProducts(c *fiber.Ctx) error {
	key := c.Params("key")
	return c.JSON(repository.FindProductsByKeyName(key))
}
