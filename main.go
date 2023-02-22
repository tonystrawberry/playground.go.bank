package main

import (
	"database/sql"
	"log"

	"github.com/tonystrawberry/playground.go.bank/api"
	db "github.com/tonystrawberry/playground.go.bank/db/sqlc"
	"github.com/tonystrawberry/playground.go.bank/util"

	_ "github.com/lib/pq"
)

func main() {
	// Load the configuration.
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Open a database connection.
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	// Connect to the database.
	store := db.NewStore(conn)

	// Create a new HTTP server and start it.
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
