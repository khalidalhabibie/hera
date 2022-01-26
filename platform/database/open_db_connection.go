package database

import "hera/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.ArticleQueries
	// *queries.BookQueries // load queries from Book model
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
		UserQueries:    &queries.UserQueries{DB: db},    // from User model
		ArticleQueries: &queries.ArticleQueries{DB: db}, // from Article model
		// BookQueries: &queries.BookQueries{DB: db}, // from Book model
	}, nil
}
