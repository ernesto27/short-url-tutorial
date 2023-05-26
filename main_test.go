package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCreateURL(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody, _ := json.Marshal(&BodyParams{Url: "https://www.google" + time.Now().String()})
	c.Request, _ = http.NewRequest(http.MethodPost, "/create-url", bytes.NewBuffer(reqBody))

	db, err := getDB()
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
		t.Error("Error connecting to database server")
	}
	CreateUrl(c, db)

	fmt.Println(w.Code)

	if w.Code != 200 {
		t.Error("Expected 200 status code, got ", w.Code)
	}
}
