package types

type Book struct {
	ID       string
	Name     string `json:"name" db:"name"`
	Author   string `json:"author" db:"author"`
	Pages    int    `json:"pages" db:"pages"`
	GenreID  int    `json:"genre_id" db:"genre_id"`
	Filename string `json:"filename" db:"filename"`
}
