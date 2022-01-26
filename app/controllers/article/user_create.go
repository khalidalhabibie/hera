package article

import (
	"hera/app/models"
	"hera/pkg/utils"
	"hera/platform/database"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateArticle(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Article struct
	articleCreateRequest := models.ArticleCreateRequest{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&articleCreateRequest); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// checking start and published
	if articleCreateRequest.StartPublishedAt.After(articleCreateRequest.EndPublishedAt) {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Please check published time",
		})

	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a article model.
	validate := utils.NewValidator()

	article := &models.Article{
		ID:               uuid.New(),
		AuthorID:         claims.UserID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		StartPublishedAt: articleCreateRequest.StartPublishedAt,
		EndPublishedAt:   articleCreateRequest.EndPublishedAt,
		Title:            articleCreateRequest.Title,
		Body:             articleCreateRequest.Body,
		IsPublished:      true,
		Category:         articleCreateRequest.Category,
	}

	// Validate article fields.
	if err := validate.Struct(article); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete book by given ID.
	if err := db.CreateArticle(article); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"article": article,
	})
}
