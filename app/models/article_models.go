package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID               uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	AuthorID         uuid.UUID `db:"author_id" json:"author_id" validate:"required,uuid"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	StartPublishedAt time.Time `db:"start_published_at" json:"start_published_at"`
	EndPublishedAt   time.Time `db:"end_published_at" json:"end_published_at"`
	Title            string    `db:"title" json:"title" validate:"required,lte=255"`
	Body             string    `db:"body" json:"body" validate:"required"`
	IsPublished      bool      `db:"is_published" json:"is_published" validate:"required"`
	Category         string    `db:"category" json:"category" validate:"required"`
}

type ArticleCreateRequest struct {
	Title            string    `db:"title" json:"title" validate:"required,lte=255"`
	Body             string    `db:"body" json:"body" validate:"required"`
	IsPublished      bool      `db:"is_published" json:"is_published" validate:"required"`
	Category         string    `db:"category" json:"category" validate:"required"`
	StartPublishedAt time.Time `db:"start_published_at" json:"start_published_at"`
	EndPublishedAt   time.Time `db:"end_published_at" json:"end_published_at"`
}

const (

	// article category
	ArticleCategorySport    = "sport"
	ArticleCategoryTech     = "tech"
	ArticleCategoryBusiness = "business"
)
