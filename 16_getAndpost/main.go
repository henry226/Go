package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func handleRequest(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	//var name = ""
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")
		insertQuery(username, password)
		fmt.Fprintf(w, "Username = %s\n", username)
		fmt.Fprintf(w, "Password = %s\n", password)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

	//fmt.Println(name)

}

func insertQuery(username string, password string) {
	//insert, err := db.Query("INSERT INTO test VALUES ( 1, '" + name + "' )")
	insert, err := db.Query("INSERT INTO test VALUES ('" + username + "', '" + password + "')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully Insert to DB")

	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to DB")

	// render webpage
	http.HandleFunc("/", handleRequest)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

	// defer the close till after the main function has finished executing
	defer db.Close()
}
