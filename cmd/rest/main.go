package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasmmo/targon-space/pkg/create"
	"github.com/lucasmmo/targon-space/pkg/database/memory"
	"github.com/lucasmmo/targon-space/pkg/get"
	"github.com/lucasmmo/targon-space/pkg/http/rest"
	"github.com/lucasmmo/targon-space/pkg/list"
	"github.com/lucasmmo/targon-space/pkg/post"
)

var (
	port, host, dbName string
	server             *gin.Engine
)

func main() {
	Init()

	repo := memory.NewRepository(make(map[string]*post.Post))
	rest.NewPostRoutes(server, get.NewUsecase(repo), create.NewUsecase(repo), list.NewUsecase(repo))

	server.Run(":" + port)
}

func Init() {
	host = environmentVar("DB_HOST")
	dbName = environmentVar("DB_NAME")
	port = environmentVar("PORT")

	server = rest.GetEngine()
}

func environmentVar(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error to load .env")
	}

	return os.Getenv(key)
}
