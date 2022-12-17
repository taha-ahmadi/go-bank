package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/taha-ahmadi/go-bank/api"
	db "github.com/taha-ahmadi/go-bank/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
