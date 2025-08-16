package main

import "database/sql"

var dsn = ""

func listing1() error {
	// sql Open doesnt always establish the connection
	// and the returning *sql.DB is a pooling connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// we could make it establish the connection and test if there is an
	// error by using the Ping method
	if err := db.Ping(); err != nil {
		return err
	}

	_ = db
	return nil
}
