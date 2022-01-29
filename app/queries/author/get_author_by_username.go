package author

import (
	"fmt"
	"hera/app/models"
	"log"
)

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
