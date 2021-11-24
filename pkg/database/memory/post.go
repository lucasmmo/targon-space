package memory

import (
	"errors"
	"strconv"

	"github.com/lucasmmo/targon-space/pkg/post"
)

type repository struct {
	db map[string]*post.Model
}

func NewRepository(db map[string]*post.Model) post.Repository {
	return &repository{db}
}

func (s *repository) Save(post *post.Model) error {
	s.db[strconv.Itoa(len(db)+1)] = post
	return nil
}

func (s *repository) FindById(id string) (post.Model, error) {
	if len(db) == 0 {
		return post.Model{}, errors.New("error")
	}
	return *s.db[id], nil
}

func (s *repository) FindAll() []post.Model {
	var posts = []post.Model{}
	for _, post := range s.db {
		posts = append(posts, *post)
	}
	return posts
}
