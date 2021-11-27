package list_test

import (
	"testing"

	"github.com/lucasmmo/targon-space/pkg/database/memory"
	"github.com/lucasmmo/targon-space/pkg/list"
)

func TestListPost(t *testing.T) {
	db := memory.GetEngine()
	repo := memory.NewRepository(db)

	t.Run("Testing Get Post", func(t *testing.T) {
		l := list.NewUsecase(repo)
		if _, err := l.Execute(); err.Error() != "no post created yet" {
			t.Error(err.Error())
		}
	})
}
