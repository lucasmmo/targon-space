package get

import (
	"errors"

	"github.com/lucasmmo/targon-space/pkg/post"
)

type usecase struct {
	repository post.Repository
}

func NewUsecase(repository post.Repository) post.Read {
	return &usecase{repository}
}

func (s *usecase) Execute(id string) (post.Post, error) {
	if id == "" {
		return post.Post{}, errors.New("invalid data")
	}

	p, err := s.repository.FindById(id)
	if err != nil {
		return post.Post{}, err
	}

	return p, nil
}
