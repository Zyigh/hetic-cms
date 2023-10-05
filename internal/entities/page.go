package entities

type Page struct {
	Title   string `db:"title"`
	Content string `db:"content"`
}
