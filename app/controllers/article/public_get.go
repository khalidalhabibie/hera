package article

import (
	"hera/app/models"
	"hera/pkg/utils"
	"hera/platform/database"

	fiber "github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	dataM, err := db.GetArticeles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// filter by data role
	data, err := utils.MarshalUsers(dataM, models.AuthPublic)
	if err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"articles": data,
	})
}
