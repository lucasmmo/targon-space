package memory

import (
	"errors"
	"strconv"

	"github.com/lucasmmo/targon-space/pkg/post"
)

type repository struct {
	db map[string]*post.Post
}

func NewRepository(db map[string]*post.Post) post.Repository {
	return &repository{db}
}

func (s *repository) Save(post *post.Post) error {
	s.db[strconv.Itoa(len(db)+1)] = post
	return nil
}

func (s *repository) FindById(id string) (post.Post, error) {
	if len(db) == 0 {
		return post.Post{}, errors.New("error")
	}
	return *s.db[id], nil
}

func (s *repository) FindAll() []post.Post {
	var posts = []post.Post{}
	for _, post := range s.db {
		posts = append(posts, *post)
	}
	return posts
}
