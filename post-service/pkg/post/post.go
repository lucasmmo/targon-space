package post

import (
	"time"
)

type Post struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAT time.Time `json:"updated_at,omitempty"`
}

func NewPost(title, content string, createdAt, updatedAt time.Time) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAT: updatedAt,
	}
}
