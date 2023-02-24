package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect_db() *sqlx.DB {
	db := sqlx.MustConnect("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	return db
}

func create_tables(db *sqlx.DB) {
	db.MustExec("DROP TABLE IF EXISTS shelves, genres, books, book_shelf")

	schema, _ := os.ReadFile("schema.sql")
	db.MustExec(string(schema))

	if mode, _ := os.LookupEnv("GIN_MODE"); mode != "release" {
		// if in debug mode, populate the database with some default values
		expr, _ := os.ReadFile("debug.sql")
		db.MustExec(string(expr))
	}
}

type Book struct {
	ID       string
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	Pages    int    `json:"pages" db:"pages"`
	GenreID  int    `db:"genre_id"`
	Filename string `db:"filename"`
}

func get_genre_name(db *sqlx.DB, id int) string {
	var genre string
	db.QueryRow("SELECT name FROM genres WHERE id=$1", id).Scan(&genre)
	return genre
}

func main() {
	db := connect_db()
	create_tables(db)

	r := gin.Default()

	r.GET("/books", func(ctx *gin.Context) {
		rows, err := db.Queryx("SELECT * FROM books")
		if err != nil {
			log.Printf("/books db.Query() error - %v", err)
			ctx.AbortWithStatus(400)
		}

		response := `<table>
			<tr>
				<td>Name</td>
				<td>Author</td>
				<td>Pages</td>
				<td>Genre</td>
			</tr>`
		var row Book

		for rows.Next() {
			err := rows.StructScan(&row)
			if err != nil {
				log.Printf("rows.StructScan() error - %v", err)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			response += fmt.Sprintf(
				"<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>",
				row.Name, row.Author, row.Pages, get_genre_name(db, row.GenreID),
			)
		}
		response += "</table>"

		ctx.Data(http.StatusOK, "text/html", []byte(response))
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
