package mysql

import (
	"log"

	"github.com/lucasmmo/targon-space/pkg/post"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func config(host, dbName string) {
	dsn := "root:my@tcp(" + host + ":3306)/" + dbName + "?parseTime=true"
	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err.Error())
	}
	conn.AutoMigrate(&post.Post{})
	db = conn
}

func GetEngine(host, dbName string) *gorm.DB {
	if db == nil {
		config(host, dbName)
		return db
	}
	return db
}
