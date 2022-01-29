package author

import (
	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries from Author model.
type UserQueries struct {
	*sqlx.DB
}

// GetUserByID query for getting one Author by given ID.
