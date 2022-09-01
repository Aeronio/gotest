package models

import (
	"Documents/gotest/postgres/09-package-organization/03_multiple-packages/config"
	"errors"
	"net/http"
	"strconv"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) //Order matters here, it will scan and assign based off of order in postgres table
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {

	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {

	//Gets the Form Values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	//This validates if every field was filled in
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	//Converts the price value
	float64Num, err := strconv.ParseFloat(p, 32) //the 32 is the bit size of the float, 32 or 64. ParseFloat always returns a Float64 but specifying 32 makes it easily converitble to a Float32
	if err != nil {
		return bk, errors.New("406. Price must be a number")
	}
	bk.Price = float32(float64Num)

	//Inserts the values into the config.DB
	_, err = config.DB.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())
	}

	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {

	//Gets the Form Values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	//This validates if every field was filled in
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	//Converts the price value
	float64Num, err := strconv.ParseFloat(p, 32) //the 32 is the bit size of the float, 32 or 64. ParseFloat always returns a Float64 but specifying 32 makes it easily converitble to a Float32
	if err != nil {
		return bk, errors.New("406. Price must be a number")
	}
	bk.Price = float32(float64Num)

	//Inserts the values into the config.DB
	_, err = config.DB.Exec("UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())
	}

	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	//Deletes a book entry
	_, err := config.DB.Exec("DELETE FROM books WHERE isbn=$1", isbn)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
