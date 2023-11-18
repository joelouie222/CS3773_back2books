package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joelouie222/CS3773_back2books/pkg"
)

//var db *sql.DB

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	fmt.Println("running main")
	http.HandleFunc("/", pkg.IndexHandler)
	http.HandleFunc("/login", pkg.LoginHandler)
	http.HandleFunc("/register", pkg.RegisterHandler)
	http.HandleFunc("/products", pkg.BooksHandler)
	http.HandleFunc("/fetch", pkg.FetchHandler)
	http.HandleFunc("/add", pkg.AddHandler)
	http.HandleFunc("/delete", pkg.DeleteHandler)

	port := getPort()
	fmt.Printf("Server started on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":80" // Default to port 8080 if PORT is not set
	} else {
		port = ":" + port // Add the colon if PORT is set
	}
	return port
}
