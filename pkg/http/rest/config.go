package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func config() {
	r := gin.Default()
	r.GET("/healthcheck", HealthCheck)
	router = r
}

func GetEngine() *gin.Engine {
	if router == nil {
		config()
		return router
	}
	return router
}

func HealthCheck(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
