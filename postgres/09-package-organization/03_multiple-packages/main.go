package main

//Multiple packages leaves a lot of room for modularity

import (
	"Documents/gotest/postgres/09-package-organization/03_multiple-packages/models"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", models.BooksIndex)
	http.HandleFunc("/books/show", models.BooksShow)
	http.HandleFunc("/books/create", models.BooksCreateForm)
	http.HandleFunc("/books/update", models.BooksUpdateForm)
	http.HandleFunc("/books/create/process", models.BooksCreateProcess)
	http.HandleFunc("/books/update/process", models.BooksUpdateProcess)
	http.HandleFunc("/books/delete/process", models.BooksDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

//TEST
