package main

import (
	"log"
	"os"

	"github.com/3elDU/book-management/api"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func load_envrionment() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("Failed to load variables from .env")
	}
}

func connect_db() *sqlx.DB {
	db := sqlx.MustConnect("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	return db
}

func create_tables(db *sqlx.DB) {
	if mode := os.Getenv("GIN_MODE"); mode != "release" {
		db.MustExec("DROP TABLE IF EXISTS shelves, genres, books, book_shelf")

		schema, _ := os.ReadFile("schema.sql")
		db.MustExec(string(schema))

		// if in debug mode, populate the database with some default values
		expr, _ := os.ReadFile("debug.sql")
		db.MustExec(string(expr))
	}
}

func main() {
	load_envrionment()

	// althorugh godotenv loads all environment variables, gin doesn't see them for some reason
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := connect_db()
	create_tables(db)

	server := api.CreateServer(db)
	server.Run()
}
