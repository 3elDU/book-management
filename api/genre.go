package api

import "github.com/gin-gonic/gin"

func (server *Server) NewGenre(ctx *gin.Context) {
	name := ctx.Param("name")
	server.db.Exec("INSERT INTO genres (name) VALUES ($1)", name)
}

func (server *Server) getGenreName(id int) string {
	var genre string
	server.db.QueryRow("SELECT name FROM genres WHERE id=$1", id).Scan(&genre)
	return genre
}
