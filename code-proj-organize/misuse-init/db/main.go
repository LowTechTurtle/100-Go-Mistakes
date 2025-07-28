package main

import (
	"database/sql"
	"net/http"
	"os"
)

// bad practice: - use global var db that anyone can read and modify
// - error handle in init func( not much we can do aside from paniking), prevent others define custom
// logic and handle the error in a better way, like fallback or retry
// - hard to test this code, init func will execute before the code, using global var
// make code hard to isolate and test properly

var db *sql.DB

func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	err = d.Ping()

	if err != nil {
		panic(err)
	}

	db = d
}

// this pattern is what we should do
// - error handling is for the user to decide, more flexible
// - create an intergration test that make sure this func works
// - connection pool is encapsulated in this function
func CreateClient(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// this init function cant fail( they can if handler is nil, but they normally dont)
// no global variable, so it doesnt affect test
func goodInitFunc() {
	redirect := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	http.HandleFunc("/blog", redirect)
	http.HandleFunc("/blog/", redirect)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/favicon.ico", static)
	http.Handle("/fonts.css", static)
	http.Handle("/fonts/", static)
	http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/",
		http.HandlerFunc(staticHandler))) // a stub imaginary handler
}
