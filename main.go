package main

import (
	"os"
	"shorturl/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type BodyParams struct {
	Url string `json:"url"`
}

func main() {
	myDB, err := getDB()
	if err != nil {
		panic(err)
	}

	defer myDB.Close()
	r := gin.Default()
	r.POST("/create-url", func(c *gin.Context) {
		CreateUrl(c, myDB)
	})

	r.GET("/:hash", func(c *gin.Context) {
		GetURL(c, myDB)
	})

	r.Run()
}

func getDB() (*db.Mysql, error) {
	user := "root"
	password := "1111"

	host := "localhost"
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}

	port := "3306"
	database := "short-url"

	return db.NewMysql(host, user, password, port, database)
}
