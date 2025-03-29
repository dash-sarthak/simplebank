package main

import (
	"database/sql"
	"log"

	"github.com/dash-sarthak/simpleBank/api"
	db "github.com/dash-sarthak/simpleBank/db/sqlc"
	"github.com/dash-sarthak/simpleBank/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@127.0.0.1:5432/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
