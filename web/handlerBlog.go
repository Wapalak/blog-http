package web

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
	"prostoTak"
	"text/template"
)

func (h *Handler) HelloPage() http.HandlerFunc {
	tmpl := template.Must(template.New("helloPage.html").ParseFiles("C:\\Users\\" +
		"User\\GolandProjects\\prostoTak\\" +
		"web\\templates\\helloPage.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w,
			"helloPage.html",
			nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) BlogList() http.HandlerFunc {
	type data struct {
		Blogs []prostoTak.Post
	}
	tmpl := template.Must(template.New("blogList.html").ParseFiles(
		"C:\\Users\\User\\GolandProjects\\" +
			"prostoTak\\web\\templates\\blogList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.blog.BlogList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data{Blogs: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) Blog() http.HandlerFunc {
	type data struct {
		Blogs prostoTak.Post
	}
	tmpl := template.Must(template.New(
		"blogDetails.html").ParseFiles(
		"C:\\Users\\User\\GolandProjects\\prostoTak\\web\\templates\\blogDetails.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		tt, err := h.blog.BlogFind(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		err = tmpl.Execute(w, tt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) BlogCreate() http.HandlerFunc {
	tmpl := template.Must(template.New(
		"blogCreate.html").ParseFiles(
		"C:\\Users\\User\\GolandProjects\\prostoTak\\web\\templates\\blogCreate.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) BlogSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.FormValue("author")
		title := r.FormValue("title")
		text := r.FormValue("text")
		//votes := 0
		if err := h.blog.BlogSave(&prostoTak.Post{
			ID:     uuid.New().String(),
			Author: author,
			Title:  title,
			Text:   text,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/blog/list", http.StatusFound)
	}
}

func (h *Handler) BlogDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		if _, err := h.blog.BlogFind(idStr); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err := h.blog.BlogDelete(idStr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/blog/list", http.StatusFound)
	}
}

func (h *Handler) BlogUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if err := h.blog.BlogUp(idStr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		redirect := fmt.Sprintf("/blog/%s", idStr)
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

func (h *Handler) BlogDown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if err := h.blog.BlogDown(idStr); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		redirect := fmt.Sprintf("/blog/%s", idStr)

		http.Redirect(w, r, redirect, http.StatusFound)
	}
}
