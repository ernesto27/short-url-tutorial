package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type BodyParams struct {
	Url string `json:"url"`
}

func main() {
	user := "root"
	password := "1111"
	host := "mysql"
	port := "3306"
	database := "short-url"

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping: " + err.Error())
		panic(err)
	}

	defer db.Close()

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

		query := "INSERT INTO short_url (hash, url) VALUES (?, ?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error inserting url",
			})
			return
		}
		defer stmt.Close()

		hasher := sha1.New()
		hasher.Write([]byte(bodyParams.Url))
		sha := hex.EncodeToString(hasher.Sum(nil))
		hash := sha[:7]

		_, err = stmt.Exec(hash, bodyParams.Url)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error inserting url",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "short url created http://localhost.8080/" + hash,
		})
	})

	r.GET("/get-url/:hash", func(c *gin.Context) {
		hash := c.Param("hash")

		query := "SELECT url FROM short_url WHERE hash = ?"
		stmt, err := db.Prepare(query)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Not found url",
			})
			return
		}
		defer stmt.Close()

		var url string
		err = stmt.QueryRow(hash).Scan(&url)
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
