package post

type Repository interface {
	Save(post *Post) error
	FindById(id string) (Post, error)
	FindAll() []Post
}
