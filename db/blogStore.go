package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"prostoTak"
)

type BlogStore struct {
	*sqlx.DB
}

func (s *BlogStore) BlogFind(id string) (prostoTak.Post, error) {
	var post prostoTak.Post
	err := s.QueryRow("SELECT id, author, title, text, votes FROM blog WHERE id = $1", id).Scan(
		&post.ID, &post.Author, &post.Title, &post.Text, &post.Votes)
	if err != nil {
		return post, fmt.Errorf(
			"error getting list of blog: %w", err)
	}
	//var result []prostoTak.Post
	//for _, post := range t {
	//	result = append(result, *post)
	//}

	return post, nil
}

func (s *BlogStore) BlogList() ([]prostoTak.Post, error) {
	var t []prostoTak.Post
	if err := s.Select(&t, "SELECT * FROM blog ORDER BY votes DESC;"); err != nil {
		return nil, fmt.Errorf(
			"error getting list of blog: %w", err)
	}
	return t, nil
}

func (s *BlogStore) BlogSave(t *prostoTak.Post) error {
	stmt, err := s.Prepare("INSERT INTO blog (author, title, text, votes) VALUES ($1, $2, $3, $4);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Author, t.Title, t.Text, t.Votes)
	if err != nil {
		log.Print("error saving blog: ", err)
		return fmt.Errorf("error saving blog: %w", err)
	}
	return nil
}

func (s *BlogStore) BlogDelete(id string) error {
	if _, err := s.Exec("DELETE FROM blog WHERE id = $1;", id); err != nil {
		return fmt.Errorf("error deleting blog: %w", err)
	}
	return nil
}

func (s *BlogStore) BlogUp(id string) error {
	if _, err := s.Exec("UPDATE blog SET votes = votes + 1 WHERE id = $1;", id); err != nil {
		return fmt.Errorf("error votting up: %w", err)
	}
	return nil
}

func (s *BlogStore) BlogDown(id string) error {
	if _, err := s.Exec("UPDATE blog SET votes = votes - 1 WHERE id = $1;", id); err != nil {
		return fmt.Errorf("error votting down: %w", err)
	}
	return nil
}
