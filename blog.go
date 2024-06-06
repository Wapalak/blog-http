package prostoTak

type Store interface {
	BlogStore
}

type Post struct {
	ID     string `db:"id"`
	Author string `db:"author"`
	Title  string `db:"title"`
	Text   string `db:"text"`
	Votes  int    `db:"votes"`
}

type BlogStore interface {
	BlogFind(id string) (Post, error)
	BlogList() ([]Post, error)
	BlogSave(t *Post) error
	BlogDelete(id string) error
	BlogUp(id string) error
	BlogDown(id string) error
}
