package main

import (
	"fmt"
	"net/http"
	"shorturl/db"

	"github.com/gin-gonic/gin"
)

func CreateUrl(c *gin.Context, myDB *db.Mysql) {
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
}

func GetURL(c *gin.Context, myDB *db.Mysql) {
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
}
