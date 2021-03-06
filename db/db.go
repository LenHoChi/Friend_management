package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize (username, password, database string) (Database, error) {
	db := Database{}

	conn, err := sql.Open("postgres", "user=postgres password=len123 dbname=social_network sslmode=disable")
	
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err!= nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
func Initialize2 () (Database, error) {
	db := Database{}

	conn, err := sql.Open("postgres", "user=postgres password=len123 dbname=social_network sslmode=disable")
	
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err!= nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}