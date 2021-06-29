package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/project/model"
	"ocg.com/project/repository"
)

func CreateNewOder(c *fiber.Ctx) error {
	order := new(model.Order)
	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(repository.CreateOder(order))
}

func GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(repository.FindOrderById(int64(id)))
}
