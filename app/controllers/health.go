package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {

	// Return status 401 and unauthorized error message.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "aman",
	})

}
