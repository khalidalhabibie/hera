package queries

import (
	"fmt"
	"hera/app/models"
	"log"

	"github.com/jmoiron/sqlx"
)

// BookQueries struct for queries from Article model.
type ArticleQueries struct {
	*sqlx.DB
}

// CreateBook method for creating book by given Article object.
func (q *ArticleQueries) CreateArticle(b *models.Article) error {
	// Define query string.
	query := `INSERT INTO articles VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.AuthorID, b.CreatedAt, b.UpdatedAt, b.StartPublishedAt, b.EndPublishedAt, b.Title, b.Body, b.IsPublished, b.Category)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

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
