package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect_db() *sqlx.DB {
	db := sqlx.MustConnect("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	return db
}

func create_tables(db *sqlx.DB) {
	if mode, _ := os.LookupEnv("GIN_MODE"); mode != "release" {
		db.MustExec("DROP TABLE IF EXISTS shelves, genres, books, book_shelf")

		schema, _ := os.ReadFile("schema.sql")
		db.MustExec(string(schema))

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
	GenreID  int    `json:"genre_id" db:"genre_id"`
	Filename string `json:"filename" db:"filename"`
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

	r.Static("/static/", "./public")

	r.GET("/books", func(ctx *gin.Context) {
		rows, err := db.Queryx("SELECT * FROM books")
		if err != nil {
			log.Printf("/books db.Query() error - %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		response := `<table>
			<tr style="font-weight: bold">
				<td>Name</td>
				<td>Author</td>
				<td>Pages</td>
				<td>Genre</td>
				<td>Filename</td>
			</tr>`
		var row Book

		for rows.Next() {
			err := rows.StructScan(&row)
			if err != nil {
				log.Printf("rows.StructScan() error - %v", err)
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
			response += fmt.Sprintf(
				"<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>",
				row.Name, row.Author, row.Pages, get_genre_name(db, row.GenreID), row.Filename,
			)
		}
		response += "</table>"

		ctx.Data(http.StatusOK, "text/html", []byte(response))
	})

	r.GET("/download/:bookid", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("bookid"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		var filename string
		db.QueryRow("SELECT filename FROM books WHERE id=$1", id).Scan(&filename)

		if filename == "" {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.File(filename)
	})

	r.GET("/shelves", func(ctx *gin.Context) {
		shelves, err := db.Queryx("SELECT name, id FROM shelves")
		if err != nil {
			log.Printf("/shelvs db.Query() error - %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		var response string

		for shelves.Next() {
			var shelf_name string
			var shelf_id int
			shelves.Scan(&shelf_name, &shelf_id)

			response += fmt.Sprintf("<h2>%v</h2> <p>", shelf_name)

			books_in_shelf, _ := db.Queryx("SELECT book_id FROM book_shelf WHERE shelf_id=$1", shelf_id)
			for books_in_shelf.Next() {
				var book_id int
				var book_name string

				books_in_shelf.Scan(&book_id)
				db.QueryRowx("SELECT name FROM books WHERE id=$1", book_id).Scan(&book_name)
				response += fmt.Sprintf("%v<br/>", book_name)
			}
			response += "</p>"
		}

		ctx.Data(http.StatusOK, "text/html", []byte(response))
	})

	r.POST("/new_book", func(ctx *gin.Context) {
		var book Book
		if err := ctx.BindJSON(&book); err != nil {
			log.Printf("/new_book ctx.BindJson() error - %v", err)
			return
		}

		_, err := db.Exec("INSERT INTO books (name, author, pages, genre_id) VALUES ($1, $2, $3, $4)",
			book.Name, book.Author, book.Pages, book.GenreID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	})

	r.POST("/upload", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		author := ctx.PostForm("author")

		file, err := ctx.FormFile("file")
		if err != nil {
			log.Printf("/upload ctx.FormFile() error - %v", err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		filename := filepath.Join("./uploads", filepath.Base(file.Filename))
		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			log.Printf("/upload ctx.SaveUploadedFile() error - %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		db.Exec("INSERT INTO books (name, author, pages, filename) VALUES ($1, $2, $3, $4)",
			name, author, -1, filename)
		ctx.String(http.StatusOK, "Successfully uploaded a file")
	})

	r.POST("/new_genre/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		db.Exec("INSERT INTO genres (name) VALUES ($1)", name)
	})

	r.Run()
}
