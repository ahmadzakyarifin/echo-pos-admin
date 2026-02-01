package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnettDB() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",user,password,host,port,dbname)

	db , err := sqlx.Connect("mysql",dsn)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak merespon: %v", err)
	}

	return db
}