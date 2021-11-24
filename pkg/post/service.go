package post

type Create interface {
	Execute(title, content string) error
}

type Read interface {
	Execute(id string) (Post, error)
}

type List interface {
	Execute() ([]Post, error)
}
