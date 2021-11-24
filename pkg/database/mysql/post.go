package mysql

import (
	"github.com/lucasmmo/targon-space/pkg/post"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) post.Repository {
	return &repository{db}
}

func (s *repository) Save(post *post.Post) error {
	if err := s.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (s *repository) FindById(id string) (post.Post, error) {
	p := post.Post{}

	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return post.Post{}, err
	}
	return p, nil
}

func (s *repository) FindAll() []post.Post {
	var posts []post.Post
	if err := s.db.Find(&posts).Error; err != nil {
		return []post.Post{}
	}
	return posts
}
