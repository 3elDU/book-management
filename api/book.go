package api

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/3elDU/book-management/types"
	"github.com/gin-gonic/gin"
)

func (server *Server) newBook(ctx *gin.Context) {
	var book types.Book
	if err := ctx.BindJSON(&book); err != nil {
		log.Printf("/new_book ctx.BindJson() error - %v", err)
		return
	}

	if _, err := server.db.Exec(
		"INSERT INTO books (name, author, pages, genre_id) VALUES ($1, $2, $3, $4)",
		book.Name, book.Author, book.Pages, book.GenreID); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func (server *Server) listBooks(ctx *gin.Context) {
	rows, err := server.db.Queryx("SELECT * FROM books")
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
	var row types.Book

	for rows.Next() {
		err := rows.StructScan(&row)
		if err != nil {
			log.Printf("rows.StructScan() error - %v", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		response += fmt.Sprintf(
			"<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>",
			row.Name, row.Author, row.Pages, server.getGenreName(row.GenreID), row.Filename,
		)
	}
	response += "</table>"

	ctx.Data(http.StatusOK, "text/html", []byte(response))
}

func (server *Server) uploadBook(ctx *gin.Context) {
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

	server.db.Exec("INSERT INTO books (name, author, pages, filename) VALUES ($1, $2, $3, $4)",
		name, author, -1, filename)
	ctx.String(http.StatusOK, "Successfully uploaded a file")
}

func (server *Server) downloadBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("bookid"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	var filename string
	server.db.QueryRow("SELECT filename FROM books WHERE id=$1", id).Scan(&filename)

	if filename == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.File(filename)
}
