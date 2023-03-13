package main

import (
	"database/sql"
	"github.com/ku-kim/learn-go-with-project-simplebank/util"

	"github.com/ku-kim/learn-go-with-project-simplebank/api"
	db "github.com/ku-kim/learn-go-with-project-simplebank/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoacConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
