package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gobank/api"
	db "gobank/db/sqlc"
	"gobank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("config error", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("caccnot start server", err.Error())
	}
	return
}
