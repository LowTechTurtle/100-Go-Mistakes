package main

import "database/sql"

func listing1(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		department string
		age        int
	)
	for rows.Next() {
		// if department is null in sql table, the string won't be ""
		// it needs to be nil and since value of string can't be nil
		// => error converting nil to string
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}

func listing2(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		// make the department a pointer to string => can handle nil
		department *string
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}

func listing3(db *sql.DB, id string) error {
	rows, err := db.Query("SELECT DEP, AGE FROM EMP WHERE ID = ?", id)
	if err != nil {
		return err
	}
	// Defer closing rows

	var (
		// or we could use NullString instead of pointers
		// the two approachs is equally good
		department sql.NullString
		age        int
	)
	for rows.Next() {
		err := rows.Scan(&department, &age)
		if err != nil {
			return err
		}
		// ...
	}
	return nil
}
