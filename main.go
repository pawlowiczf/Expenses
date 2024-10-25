package main

import (
	"database/sql"
	"expenses/api"
	"expenses/config"
	db "expenses/db/sqlc"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config file: %s", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot open db: %s", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatalf("cannot ping db: %s", err)
	}
	
	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatalf("cannot create new HTTP server: %s", err)
	}

	err = server.RunServer(config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot run HTTP server: %s", err)
	}
}
