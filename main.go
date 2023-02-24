package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func connect_db() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	log.Println(db, err)
	return db
}

func create_table(db *sql.DB) {
	res, err := db.Exec("DROP TABLE IF EXISTS books")
	log.Println(res, err)

	res, err = db.Exec(`CREATE TABLE books (
		id SERIAL PRIMARY KEY,
		name VARCHAR(128) NOT NULL,
		author VARCHAR(64) NOT NULL,
		pages INT NOT NULL,
		filename TEXT
	);`)
	log.Println(res, err)
}

type Book struct {
	ID       string
	Name     string `json:"name"`
	Author   string `json:"author"`
	Pages    int    `json:"pages"`
	Filename string
}

func main() {
	db := connect_db()
	create_table(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})

	r.POST("/new_book", func(ctx *gin.Context) {
		var book Book
		if err := ctx.BindJSON(&book); err != nil {
			log.Printf("/new_book ctx.BindJson() error - %v", err)
			return
		}

		res, err := db.Exec("INSERT INTO books (name, author, pages) VALUES ($1, $2, $3)", book.Name, book.Author, book.Pages)
		log.Println(res, err)
	})

	r.Run()
}
