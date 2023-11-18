package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//var db *sql.DB

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	fmt.Println("running main")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/products", BooksHandler)
	http.HandleFunc("/fetch", FetchHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/delete", DeleteHandler)

	port := getPort()
	fmt.Printf("Server started on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Default to port 8080 if PORT is not set
	}
	return port
}
