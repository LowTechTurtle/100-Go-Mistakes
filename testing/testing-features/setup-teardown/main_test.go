package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestMySQLIntegration(t *testing.T) {
	// setup for each test function( per test)
	setupMySQL()
	defer teardownMySQL()

	// ...
}

func createConnection(t *testing.T, dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.FailNow()
	}
	// teardown for only one of the helper function
	// make other test func easier for they no longer need to rewrite this in the future:w
	t.Cleanup(
		func() {
			_ = db.Close()
		})
	return db
}

func TestMain(m *testing.M) {
	// setup for the package
	// every tests will have MySQL already setup and it will close after finishing the tests
	setupMySQL()
	code := m.Run()
	teardownMySQL()
	os.Exit(code)
}

func setupMySQL() {}

func teardownMySQL() {}
