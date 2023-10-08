package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gobank/util"
	"log"
	"os"
	"testing"
)

var testQuaries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conf, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("config load error ", err)
	}
	testDB, err = sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQuaries = New(testDB)
	os.Exit(m.Run())
}
