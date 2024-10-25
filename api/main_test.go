package api

import (
	"database/sql"
	"expenses/config"
	db "expenses/db/sqlc"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

var serverTest *Server 

func TestMain(m *testing.M) {
	//
	config, err := config.LoadConfig("..")
	if err != nil {
		log.Fatalf("cannot load config file: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}

	store := db.NewStore(conn)
	serverTest = TestServer(store, config)

	os.Exit(m.Run())
}
