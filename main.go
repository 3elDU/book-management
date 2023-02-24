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
		name VARCHAR(128),
		author VARCHAR(64),
		pages INT,
		filename TEXT
	);`)
	log.Println(res, err)
}

func main() {
	db := connect_db()
	create_table(db)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})
	r.Run()
}
