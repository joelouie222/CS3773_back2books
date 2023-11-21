package pkg

import (
	"database/sql"
	"fmt"
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
func createBook(Title string, Prod string, ISBN string, Date string, Price float32, Format string, Publisher string) error {
	_, err := db.Exec(
		"INSERT INTO BOOKS (book_title, prod_desc, book_ISBN, BOOK_PUBLISHED_DATE, PRICE, BOOK_FORMAT, NUM_PAGES, PUBLISHER_NAME) values (?, ?, ?, ?, ?, ?, ?, ?)",
		Title, Prod, ISBN, Date, Price, Format, Publisher)

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
	rows, err := db.Query("SELECT b.*, bi.IMAGE_LINK FROM BOOKS b LEFT JOIN BOOK_IMAGE bi ON b.BOOK_ID = bi.book_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Desc, &book.ISBN, &book.PublishDate, &book.Price, &book.Format, &book.NumPages, &book.Publisher, &book.ImgSrc)
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

// Returns specfic item from the BOOKS table.
func getBookByID(id int) (Book, error) {
	var book Book
	row, err := db.Query("SELECT b.*, bi.IMAGE_LINK FROM BOOKS b LEFT JOIN BOOK_IMAGE bi ON b.BOOK_ID = bi.book_id WHERE b.BOOK_ID=?", id)
	if err != nil {
		return book, err
	}
	defer row.Close()
	if row.Next() {
		err := row.Scan(&book.ID, &book.Title, &book.Desc, &book.ISBN, &book.PublishDate, &book.Price, &book.Format, &book.NumPages, &book.Publisher, &book.ImgSrc)
		if err != nil {
			return book, err
		}
	} else {
		// No matching record found
		return book, fmt.Errorf("Book with ID %d not found", id)
	}

	return book, nil
}
