package auth

import (
	"log"
	"time"

	"hera/app/models"
	"hera/pkg/utils"
	"hera/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthorSignUp(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signUp := &models.SignUp{}

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		log.Println("Error Parse Sign Up Body")
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Author model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		log.Println("Error Validate Sign Up Body , error ", err)
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// check password and re password are the same
	if signUp.Password != signUp.Repassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Password and confirm password does not match",
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

	// Create a new user struct.
	author := &models.Author{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Username:     signUp.Username,
		PasswordHash: utils.GeneratePassword(signUp.Password),
		Active:       true,
	}

	// Validate user fields.
	if err := validate.Struct(author); err != nil {
		log.Println("Error Validate Check Author Body")
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// filter by data by auth
	data, err := utils.MarshalUsers(author, models.AuthAuthor)
	if err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new user with validated data.
	if err := db.CreateUser(author); err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"author": data,
	})
}
