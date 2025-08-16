package main

import "database/sql"

func listing1(db *sql.DB, id string) error {
	// for queries that repeat multiple time, instead of calling query
	// we call prepare to precompile part of the code for efficency
	stmt, err := db.Prepare("SELECT * FROM ORDER WHERE ID = ?")
	if err != nil {
		return err
	}
	// we can call the query method on the returned *sql.Stmt to query
	rows, err := stmt.Query(id)
	if err != nil {
		return err
	}
	_ = rows
	return nil
}
