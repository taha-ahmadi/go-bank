package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/taha-ahmadi/go-bank/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot read env file:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
