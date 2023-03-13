package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

/*
https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
m *testing.M는 beforeAll, afterAll 이라고 생각하면 좋음 (m.Run 기준
*/
func TestMain(m *testing.M) {
	log.Println("--- Before Test ---")
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	code := m.Run()
	log.Println("--- AfterAll Test ---")
	os.Exit(code)

}
