package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"shorturl/db"
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

	err = initDBData(db)
	if err != nil {
		fmt.Println(err)
		t.Error("Error creating table")
	}

	CreateUrl(c, db)

	fmt.Println(w.Code)

	if w.Code != 200 {
		t.Error("Expected 200 status code, got ", w.Code)
	}
}

func initDBData(myDB *db.Mysql) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS short_url (
		id INT NOT NULL AUTO_INCREMENT,
		hash VARCHAR(7) NOT NULL,
		url VARCHAR(255) NOT NULL,
		created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		UNIQUE KEY (hash)
	  );
	`

	_, err := myDB.Db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}
