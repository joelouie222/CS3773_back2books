package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
)

var db *sql.DB

func main() {
    // http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

    fmt.Println("running main")
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/pages/products.html", booksHandler)
    http.HandleFunc("/fetch", fetchHandler)
    http.HandleFunc("/add", AddHandler)
    http.HandleFunc("/delete", DeleteHandler)

    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
