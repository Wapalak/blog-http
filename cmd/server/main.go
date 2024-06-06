package main

import (
	"log"
	"net/http"
	"prostoTak/db"
	"prostoTak/web"
)

func main() {
	blog, err := db.NewStore("postgres://postgres:" +
		"12345@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(blog)
	http.ListenAndServe(":8080", h)
}
