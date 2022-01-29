package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe Author object.
type Author struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid" groups:"author"`
	CreatedAt    time.Time `db:"created_at" json:"created_at" groups:"author"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at" groups:"author"`
	PasswordHash string    `db:"password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	Username     string    `db:"username" json:"username" validate:"required,lte=255" groups:"author"`
	Active       bool      `db:"active" json:"active" validate:"required" groups:"author"`
}
