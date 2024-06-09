package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgre() *sql.DB {
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	dbname := os.Getenv("DBName")

	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal("DB open fail", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping fail", err.Error())
	}

	return db
}
