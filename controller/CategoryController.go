package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/project/repository"
)

func GetAllCategories(c *fiber.Ctx) error {
	return c.JSON(repository.FindAllCategories())
}
