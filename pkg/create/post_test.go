package create_test

import (
	"testing"

	"github.com/lucasmmo/targon-space/pkg/create"
	"github.com/lucasmmo/targon-space/pkg/database/memory"
)

func TestCreatePost(t *testing.T) {
	db := memory.GetEngine()
	repo := memory.NewRepository(db)

	t.Run("Testing Create Post", func(t *testing.T) {
		s := create.NewUsecase(repo)
		if err := s.Execute("test", "test"); err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("Testing Create Post Without Title and/or Content", func(t *testing.T) {
		s := create.NewUsecase(repo)
		if err := s.Execute("", ""); err == nil {
			t.Error(err.Error())
		}
	})
}
