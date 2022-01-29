package article

import (
	"fmt"
	"hera/app/models"
	"log"
)

// GetUserByID query for getting one Articel by given ID.
func (q *ArticleQueries) GetArticelByID(id string) (*models.Article, error) {
	// Define Author variable.
	article := &models.Article{}

	// Define query string.
	query := `SELECT * FROM "articles" WHERE id = $1 `

	// Send query to database.
	err := q.Get(article, query, id)
	if err != nil {
		// Return empty object and error.
		log.Println(fmt.Sprintf("error to find %v , err : %v", id, err))
		return nil, err
	}

	// Return query result.
	return article, nil
}
