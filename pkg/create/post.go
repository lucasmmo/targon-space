package create

import (
	"errors"
	"time"

	"github.com/lucasmmo/targon-space/pkg/post"
)

type usecase struct {
	repository post.Repository
}

func NewUsecase(repo post.Repository) post.Create {
	return &usecase{repo}
}

func (s *usecase) Execute(title, content string) error {
	if title == "" || content == "" {
		return errors.New("empty data to create post")
	}

	newPost := post.NewModel(title, content, time.Now(), time.Now())
	if err := s.repository.Save(newPost); err != nil {
		return err
	}

	return nil
}
