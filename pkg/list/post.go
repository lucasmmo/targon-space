package list

import (
	"errors"

	"github.com/lucasmmo/targon-space/pkg/post"
)

type usecase struct {
	repository post.Repository
}

func NewUsecase(repository post.Repository) post.List {
	return &usecase{repository}
}

func (s *usecase) Execute() ([]post.Model, error) {
	if len(s.repository.FindAll()) == 0 {
		return []post.Model{}, errors.New("no post created yet")
	}
	return s.repository.FindAll(), nil
}
