package article

import (
	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Article model.
type ArticleQueries struct {
	*sqlx.DB
}
