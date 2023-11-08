package handlers

import (
	"github.com/lnds/lab-ic-ms/api/database"
	"github.com/lnds/lab-ic-ms/api/models"

	"github.com/gofiber/fiber/v2"
)

func ListDirectors(c *fiber.Ctx) error {
	directors := []models.Director{}
	database.DB.Db.Find(&directors)
	return c.Status(fiber.StatusOK).JSON(directors)
}

func GetDirector(c *fiber.Ctx) error {
	var director models.Director
	id := c.Params("id")
	if err := database.DB.Db.First(&director, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(director)
}

func CreateDirector(c *fiber.Ctx) error {
	director := new(models.Director)
	if err := c.BodyParser(director); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Create(director).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(director)
}

func UpdateDirector(c *fiber.Ctx) error {
	director := new(models.Director)
	if err := c.BodyParser(director); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Save(director).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(director)
}
