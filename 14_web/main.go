package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Index Page</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server Starting... localhost:3000")
	http.ListenAndServe(":3000", nil)
}
