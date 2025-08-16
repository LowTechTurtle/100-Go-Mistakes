package main

import (
	"context"
	"database/sql"
	"log"
)

func get1(ctx context.Context, db *sql.DB, id string) (string, int, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	// handle error from query
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// handle error from closing the rows
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		// handle error after scanning
		if err != nil {
			return "", 0, err
		}
	}

	return department, age, nil
}

func get2(ctx context.Context, db *sql.DB, id string) (string, int, error) {
	rows, err := db.QueryContext(ctx,
		"SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v\n", err)
		}
	}()

	var (
		department string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return "", 0, err
		}
	}
	// gotta check if an error occured when we running rows.Next()
	// or it is exited because there are more rows
	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	return department, age, nil
}
