package post

import (
	"time"
)

type Model struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAT time.Time `json:"updated_at,omitempty"`
}

func NewModel(title, content string, createdAt, updatedAt time.Time) *Model {
	return &Model{
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAT: updatedAt,
	}
}
