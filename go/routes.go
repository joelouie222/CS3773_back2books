package main

import (
    "fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)
func indexHandler(w http.ResponseWriter, r *http.Request) {

    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        log.Fatal(err)
    }
    tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var user User 
    user.username = "admin"
    user.password = "password"

    tmpl := template.Must(template.ParseFiles("./pages/login.html"))
    err := tmpl.Execute(w, nil)
    if err != nil { log.Fatal(err) }
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Register")
    tmpl, err := template.ParseFiles("pages/login.html")
    err = tmpl.Execute(w, nil)
    if err != nil { log.Fatal(err) }

}
func booksHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("bookHandling")
    var books []Book
    books, err := getBooks()
    if err != nil {
        log.Fatal(err)
    }
    tmpl, err := template.ParseFiles("pages/products.html")
    if err != nil {
        fmt.Println(err);
        log.Fatal(err)
    }

    tmpl.Execute(w, books)
}
func fetchHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("fetchHandling");
    books, err := getBooks()
    if err != nil { log.Fatal(err) }

    tmpl := template.Must(template.ParseFiles("./pages/productItem.html"))

    err = tmpl.Execute(w, books)
    if err != nil {
        log.Fatal(err)
        fmt.Println(err);
    }
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    lname := r.FormValue("lname")
    fname := r.FormValue("fname")
    price := r.FormValue("price")
    publisher := r.FormValue("publisher")

    newPrice, err := strconv.ParseFloat(price, 32)
    if err != nil {
        log.Fatal(err)
    }
    float := float32(newPrice)

    createBook(title, lname, fname, float, publisher)
    http.Redirect(w, r, "/fetch", http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.FormValue("id"))
    if err != nil {
        log.Fatal(err)
    }
    deleteBook(id)
    http.Redirect(w, r, "/fetch", http.StatusSeeOther)
}
