package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID
	Title     string
	Content   string
	AuthorID  uuid.UUID
	CreatedAt time.Time
}
