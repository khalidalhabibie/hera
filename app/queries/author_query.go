package queries

import (
	"fmt"
	"hera/app/models"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries from Author model.
type UserQueries struct {
	*sqlx.DB
}

// GetUserByID query for getting one Author by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (*models.Author, error) {
	// Define Author variable.
	user := models.Author{}

	// Define query string.
	query := `SELECT * FROM authors WHERE id = $1`

	// Send query to database.
	err := q.Get(&user, query, id)
	if err != nil {
		// Return empty object and error.

		return nil, err
	}

	// Return query result.
	return &user, nil
}

// GetUserByUsername query for getting one Author by given username.
func (q *UserQueries) GetUserByUsername(username string) (*models.Author, error) {
	// Define Author variable.
	user := &models.Author{}

	// Define query string.
	query := `SELECT * FROM "authors" WHERE username = $1 `

	// Send query to database.
	err := q.Get(user, query, username)
	if err != nil {
		log.Println(fmt.Sprintf("error to find %v , err : %v", username, err))
		// Return empty object and error.
		return nil, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given username and password hash.
func (q *UserQueries) CreateUser(u *models.Author) error {
	// Define query string.
	query := `INSERT INTO authors VALUES ($1, $2, $3, $4, $5, $6)`

	// Send query to database.
	_, err := q.Exec(
		query,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Username,
		u.PasswordHash,
		u.Active,
	)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
