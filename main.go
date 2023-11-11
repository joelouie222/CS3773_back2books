package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/joelouie222/CS3773_back2books/pkg"
)


func main() {
    // http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

    http.HandleFunc("/", pkg.IndexHandler)
    http.HandleFunc("/login", pkg.LoginHandler)
    http.HandleFunc("/register", pkg.RegisterHandler)
    http.HandleFunc("/pages/products.html", pkg.BooksHandler)
    http.HandleFunc("/fetch", pkg.FetchHandler)
    http.HandleFunc("/add", pkg.AddHandler)
    http.HandleFunc("/delete", pkg.DeleteHandler)

    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
