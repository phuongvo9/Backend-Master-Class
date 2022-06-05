package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

// All unit tests will need a object "Queries" => declare as a global variable
// {Queries struct} contains DBTX which can be a database connection or a database transaction.
// In these test cases, we will use testQueries as a db connection and use it to create a Queries object
var testQueries *Queries

const (
	dbDrive  = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable"
)

// TestMain will be the entry point for executing all unit-tests in a same GOlang package
func TestMain(t *testing.M) {
	// use sql.Open to open a database Connection
	conn, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		log.Fatalln("can not connect to the database")
	}
	// we use the connection to create a new Queries object
	testQueries = New(conn)

	// use m.Run() to start running the Unit tests and return correct exit code by os.Exit()
	os.Exit(m.Run())
}
