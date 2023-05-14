package main

import (
	"fmt"
	"net/http"
	"shorturl/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type BodyParams struct {
	Url string `json:"url"`
}

func main() {
	// TODO : Move to env
	user := "root"
	password := "1111"
	host := "mysql"
	port := "3306"
	database := "short-url"

	myDB, err := db.NewMysql(host, user, password, port, database)
	if err != nil {
		panic(err)
	}

	defer myDB.Close()

	r := gin.Default()
	r.POST("/create-url", func(c *gin.Context) {
		// TODO: Validation data
		var bodyParams BodyParams
		if err := c.ShouldBindJSON(&bodyParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error parsing body json",
			})
			return
		}

		hash, err := myDB.CreateShortURL(bodyParams.Url)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error inserting url",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "short url created http://localhost/" + hash,
		})
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		url, err := myDB.GetShortURL(hash)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not found url",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": url,
		})
	})

	r.Run()
}
