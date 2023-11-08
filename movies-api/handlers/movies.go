package handlers

import (
	"github.com/lnds/lab-ic-ms/api/database"
	"github.com/lnds/lab-ic-ms/api/models"

	"github.com/gofiber/fiber/v2"
)

func ListMovies(c *fiber.Ctx) error {
	movies := []models.Movie{}
	database.DB.Db.Preload("Director").Find(&movies)
	return c.Status(fiber.StatusOK).JSON(movies)
}

func GetMovie(c *fiber.Ctx) error {
	var movie models.Movie
	id := c.Params("id")
	if err := database.DB.Db.Preload("Director").First(&movie, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(movie)
}

func CreateMovie(c *fiber.Ctx) error {
	movie := new(models.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Create(movie).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	movie := new(models.Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Db.Save(movie).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(movie)
}
