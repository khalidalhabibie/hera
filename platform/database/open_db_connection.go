package database

import (
	"hera/app/queries/article"
	// queries "hera/app/queries/article"
	"hera/app/queries/author"
)

// Queries struct for collect all app queries.
type Queries struct {
	*author.UserQueries // load queries from User model
	*article.ArticleQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:    &author.UserQueries{DB: db},     // from User model
		ArticleQueries: &article.ArticleQueries{DB: db}, // from Article model
	}, nil
}
