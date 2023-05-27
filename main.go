package main

import (
	"os"
	"shorturl/cache"
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

	myCache := getCache()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": os.Getenv("VERSION"),
		})
	})

	r.POST("/create-url", func(c *gin.Context) {
		CreateUrl(c, myDB)
	})

	r.GET("/:hash", func(c *gin.Context) {
		GetURL(c, myDB, myCache)
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

func getCache() *cache.Redis {
	host := "localhost"
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}

	return cache.NewRedis(host)
}
