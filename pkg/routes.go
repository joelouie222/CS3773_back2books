package pkg

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	user.username = "admin"
	user.password = "password"

	tmpl := template.Must(template.ParseFiles("pages/login.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/login.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book
	books, err := getBooks()
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles("pages/products.html")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	tmpl.Execute(w, books)
}
func FetchHandler(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks()
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles("pages/productItem.html"))

	err = tmpl.Execute(w, books)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
}

func FetchBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("fetch book handler called")
	bookID := r.URL.Query().Get("id")
	if bookID == "" {
		http.Error(w, "Book ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := getBookByID(id)
	if err != nil {
		http.Error(w, "Error fetching book", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("pages/itemInfo.html"))
	err = tmpl.Execute(w, book)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func SellingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/selling.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func PromotionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/promotion.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/about.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/contact.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/favorites.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/cart.html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// func AddHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.FormValue("title")
// 	lname := r.FormValue("lname")
// 	fname := r.FormValue("fname")
// 	price := r.FormValue("price")
// 	publisher := r.FormValue("publisher")

// 	newPrice, err := strconv.ParseFloat(price, 32)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	float := float32(newPrice)

// 	createBook(title, lname, fname, float, publisher)
// 	http.Redirect(w, r, "/fetch", http.StatusSeeOther)
// }

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	deleteBook(id)
	http.Redirect(w, r, "/fetch", http.StatusSeeOther)
}
