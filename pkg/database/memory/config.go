package memory

import (
	"github.com/lucasmmo/targon-space/pkg/post"
)

var db map[string]*post.Model

func config() {
	db = map[string]*post.Model{}
}

func GetEngine() map[string]*post.Model {
	if db == nil {
		config()
		return db
	}
	return db
}
