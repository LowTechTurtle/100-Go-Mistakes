package main

import (
	"database/sql"
	"log"
)

func listing1() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM CUSTOMERS")
	if err != nil {
		return err
	}

	// Use rows
	_ = rows

	return nil
}

func listing2() error {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM CUSTOMERS")
	if err != nil {
		return err
	}

	// close and handle the error of the returned rows
	// forget to close rows lead to connection leak
	// which prevent putting back the connection to the connection pool
	//
	// *sql.DB represent a pool of connection, we rarely need to close it
	// but if we're done with the database and no application need to connect
	// to it, we should close it, and it should be long-lived and shared
	// by many goroutines
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	// Use rows
	_ = rows

	return nil
}

var dataSourceName = ""
