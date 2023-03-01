package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) listNamespaces(ctx *gin.Context) {
	shelves, err := server.db.Queryx("SELECT name, id FROM shelves")
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

		books_in_shelf, _ := server.db.Queryx("SELECT book_id FROM book_shelf WHERE shelf_id=$1", shelf_id)
		for books_in_shelf.Next() {
			var book_id int
			var book_name string

			books_in_shelf.Scan(&book_id)
			server.db.QueryRowx("SELECT name FROM books WHERE id=$1", book_id).Scan(&book_name)
			response += fmt.Sprintf("%v<br/>", book_name)
		}
		response += "</p>"
	}

	ctx.Data(http.StatusOK, "text/html", []byte(response))
}
