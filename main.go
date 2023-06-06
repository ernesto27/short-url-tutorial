package main

import (
	"fmt"
	"os"
	"shorturl/cache"
	"shorturl/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type BodyParams struct {
	Url string `json:"url"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	myDB, err := getDB()
	if err != nil {
		panic(err)
	}
	defer myDB.Close()

	myCache := getCache()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version from githubactions": os.Getenv("VERSION"),
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
	return db.NewMysql(os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
}

func getCache() *cache.Redis {
	return cache.NewRedis(os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PASSWORD"), os.Getenv("CACHE_PORT"))
}
