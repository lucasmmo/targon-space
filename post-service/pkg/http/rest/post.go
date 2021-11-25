package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasmmo/targon-space/pkg/post"
)

type postRoutes struct {
	r      *gin.Engine
	read   post.Read
	create post.Create
	list   post.List
}

func NewPostRoutes(r *gin.Engine, read post.Read, create post.Create, list post.List) *postRoutes {
	p := &postRoutes{r, read, create, list}
	p.SetupPostRoutes(p.r)
	return p
}

func (s *postRoutes) SetupPostRoutes(r *gin.Engine) {
	postRoutes := r.Group("/post")
	postRoutes.GET("/:id", s.getPost)
	postRoutes.GET("/", s.getPosts)
	postRoutes.POST("/", s.createPost)
}

func (s *postRoutes) getPost(c *gin.Context) {
	id := c.Param("id")

	post, err := s.read.Execute(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *postRoutes) createPost(c *gin.Context) {
	post := post.Post{}

	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := s.create.Execute(post.Title, post.Content); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (s *postRoutes) getPosts(c *gin.Context) {

	posts, err := s.list.Execute()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}
