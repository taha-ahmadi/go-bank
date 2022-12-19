package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/taha-ahmadi/go-bank/api"
	db "github.com/taha-ahmadi/go-bank/db/sqlc"
	"github.com/taha-ahmadi/go-bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot read env file:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("cannot run the server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
