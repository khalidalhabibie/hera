package article

import (
	"hera/platform/database"

	fiber "github.com/gofiber/fiber/v2"
)

func GetArticleByID(c *fiber.Ctx) error {

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Define articel ID.
	articelID := c.Params("id")
	articelM, err := db.GetArticelByID(articelID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"article": articelM,
	})
}
