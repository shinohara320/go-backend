package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shinohara320/travel-agent/database"
	"github.com/shinohara320/travel-agent/models"
)

func PostTestimonials(c *fiber.Ctx) error {
	var testi models.Testimonials
	if err := c.BodyParser(&testi); err != nil {
		fmt.Println("Unable to parse body")
	}
	if err := database.DB.Create(&testi).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulations, your message posted",
	})
}

func GetTestimonials(c *fiber.Ctx) error {
	var testimonials []models.Testimonials
	if err := database.DB.Find(&testimonials).Error; err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(testimonials)
}
