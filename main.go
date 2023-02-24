package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect_db() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	log.Println(db, err)
	return db
}

func create_table(db *sqlx.DB) {
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
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	Pages    int    `json:"pages" db:"pages"`
	Filename string `db:"filename"`
}

func main() {
	db := connect_db()
	create_table(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})

	r.GET("/books", func(ctx *gin.Context) {
		rows, err := db.Queryx("SELECT name, author, pages FROM books")
		if err != nil {
			log.Printf("/books db.Query() error - %v", err)
			ctx.AbortWithStatus(400)
		}

		response := `<table>
			<tr>
				<td>Name</td>
				<td>Author</td>
				<td>Pages</td>
			</tr>`
		var row Book

		for rows.Next() {
			err := rows.StructScan(&row)
			if err != nil {
				log.Printf("rows.StructScan() error - %v", err)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			response += fmt.Sprintf(
				"<tr><td>%v</td><td>%v</td><td>%v</td></tr>",
				row.Name, row.Author, row.Pages,
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
