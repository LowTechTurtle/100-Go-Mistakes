package main

import (
	"database/sql"
	"log"
)

const query = "..."

func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	// this will just ignore error
	// we could write a closure to explicitly state that
	// ignoring error is intended.
	defer rows.Close()

	// use rows
	return 0, nil
}

func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	// we should handle error like so, at the very least, we should log the error
	// or prioritize returning it when the outer err is nil
	defer func() {
		errClose := rows.Close()
		if err != nil {
			if errClose != nil {
				log.Printf("Failed to close rows: %v", errClose)
			}
			return
		}
		err = errClose
	}()

	return 0, nil
}
