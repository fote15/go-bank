package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gobank/api"
	db "gobank/db/sqlc"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://admin1:admin2@localhost:5432/gobank?sslmode=disable"
	dbAdress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(dbAdress)
	if err != nil {
		log.Fatal("caccnot start server", err.Error())
	}
	return
}
