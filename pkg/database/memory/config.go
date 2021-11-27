package memory

import (
	"github.com/lucasmmo/targon-space/pkg/post"
)

var db map[string]*post.Post

func config() {
	db = map[string]*post.Post{}
}

func GetEngine() map[string]*post.Post {
	if db == nil {
		config()
		return db
	}
	return db
}
