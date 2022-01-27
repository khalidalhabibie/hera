package article

import (
	"fmt"
	"hera/app/models"
	"log"
)

func (q *ArticleQueries) GetArticeles() ([]models.Article, error) {
	// Define Author variable.
	article := []models.Article{}

	// Define query string.
	query := `SELECT * FROM "articles" WHERE is_published = true AND start_published_at < NOW() AND end_published_at > NOW() `

	// Send query to database.
	err := q.Select(&article, query)
	if err != nil {
		// Return empty object and error.
		log.Println(fmt.Sprintf("error to find articles , err : %v", err))
		return nil, err
	}

	// Return query result.
	return article, nil
}
