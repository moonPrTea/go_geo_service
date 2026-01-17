package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	DatabaseURL string 
}

func Init(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
