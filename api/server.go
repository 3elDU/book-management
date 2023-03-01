package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db     *sqlx.DB
	router *gin.Engine
}

func CreateServer(database *sqlx.DB) *Server {
	s := &Server{
		db:     database,
		router: gin.Default(),
	}

	s.register_routes()
	return s
}

func (server *Server) register_routes() {
	server.router.Static("/static", "./public/")

	server.router.GET("/books", server.listBooks)
	server.router.GET("/download/:book", server.downloadBook)
	server.router.POST("/new_book", server.newBook)
	server.router.POST("/upload_book", server.uploadBook)

	server.router.GET("/namespaces", server.listNamespaces)
}

func (server *Server) Run() {
	server.router.Run()
}
