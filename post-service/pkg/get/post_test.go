package get_test

import (
	"testing"

	"github.com/lucasmmo/targon-space/pkg/create"
	"github.com/lucasmmo/targon-space/pkg/database/memory"
	"github.com/lucasmmo/targon-space/pkg/get"
)

func TestGetPost(t *testing.T) {
	db := memory.GetEngine()
	repo := memory.NewRepository(db)

	t.Run("Testing Get Post", func(t *testing.T) {
		g := get.NewUsecase(repo)

		c := create.NewUsecase(repo)
		if err := c.Execute("test", "test"); err != nil {
			t.Error(err.Error())
		}

		if _, err := g.Execute("1"); err != nil {
			t.Error(err.Error())
		}
	})
}
