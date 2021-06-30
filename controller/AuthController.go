package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/project/database"
	"ocg.com/project/model"
	"ocg.com/project/util"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	for key := range data {
		if data[key] == "" {
			c.Status(400)
			return c.JSON(fiber.Map{
				"message": "Cannot be null",
				"field":   key,
			})
		}
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := model.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Username:  data["username"],
		Email:     data["email"],
	}
	user.SetPassword(data["password"])

	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user model.User
	database.DB.Where("username = ?", data["username"]).First(&user)
	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(cookie)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var user model.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}
