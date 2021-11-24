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

func (s *repository) Save(post *post.Model) error {
	if err := s.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (s *repository) FindById(id string) (post.Model, error) {
	p := post.Model{}

	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return post.Model{}, err
	}
	return p, nil
}
