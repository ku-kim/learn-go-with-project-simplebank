package db

import (
	"database/sql"
	"github.com/ku-kim/learn-go-with-project-simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

/*
https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
m *testing.M는 beforeAll, afterAll 이라고 생각하면 좋음 (m.Run 기준
*/
func TestMain(m *testing.M) {
	config, err := util.LoacConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	log.Println("--- Before Test ---")
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	code := m.Run()
	log.Println("--- AfterAll Test ---")
	os.Exit(code)

}
