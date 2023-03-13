package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

/*
https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
m *testing.M는 beforeAll, afterAll이라고 생각하면 좋음 (m.Run 기준
*/
func TestMain(m *testing.M) {
	log.Println("--- Before Test ---")
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	testQueries = New(conn)
	code := m.Run()
	log.Println("--- AfterAll Test ---")
	os.Exit(code)
}
