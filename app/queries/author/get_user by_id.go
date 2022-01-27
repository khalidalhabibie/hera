package author

import (
	"hera/app/models"

	"github.com/google/uuid"
)

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
