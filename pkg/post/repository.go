package post

type Repository interface {
	Save(post *Model) error
	FindById(id string) (Model, error)
	FindAll() []Model
}
