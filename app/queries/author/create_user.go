package author

import "hera/app/models"

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
