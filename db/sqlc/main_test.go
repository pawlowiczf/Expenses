package db

import (
	"database/sql"
	"expenses/config"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

var queriesTest *Queries

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queriesTest = New(testDB)
	os.Exit(m.Run())
}
