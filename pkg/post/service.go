package post

type Create interface {
	Execute(title, content string) error
}

type Read interface {
	Execute(id string) (Model, error)
}

type List interface {
	Execute() ([]Model, error)
}
