package article

import "hera/app/models"

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
