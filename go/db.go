package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "back2books.xyz:3306",
		DBName: "back2books",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Inserts a book into the BOOKS table.
func createBook(Title string, Lname string, Fname string, Price float32, Publisher string) error {
	_, err := db.Exec(
		"INSERT INTO BOOKS (title, author_lname, author_fname, price, publisher) values (?, ?, ?, ?, ?)",
		Title, Lname, Fname, Price, Publisher)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// Deletes a book from the BOOKS table given an id.
func deleteBook(id int) error {
	_, err := db.Exec("DELETE FROM BOOKS WHERE id=?", id)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// Returns all items in the BOOKS table.
func getBooks() ([]Book, error) {
	rows, err := db.Query("SELECT * FROM BOOKS")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Desc, &book.ISBN, &book.PublishDate, &book.Price, &book.Genre, &book.Format, &book.NumPages, &book.Publisher)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return books, nil
}
