package db

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Db *sql.DB
}

func NewMysql(host, user, password, port, database string) (*Mysql, error) {
	tls := "?tls=true"
	if host == "localhost" {
		tls = ""
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", user, password, host, database, tls)
	fmt.Println("DNS", dns)

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, errors.New("error connecting to the database")
	}

	return &Mysql{
		Db: db,
	}, nil
}

func (m *Mysql) CreateShortURL(url string) (string, error) {
	query := "INSERT INTO short_url (hash, url) VALUES (?, ?)"
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	hasher := sha1.New()
	hasher.Write([]byte(url))
	sha := hex.EncodeToString(hasher.Sum(nil))
	hash := sha[:7]

	_, err = stmt.Exec(hash, url)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func (m *Mysql) GetShortURL(hash string) (string, error) {
	query := "SELECT url FROM short_url WHERE hash = ?"
	stmt, err := m.Db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var url string
	err = stmt.QueryRow(hash).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (m *Mysql) Close() {
	m.Db.Close()
}
