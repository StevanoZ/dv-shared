package shrd_utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(dbDriver, dbSource string) *sql.DB {
	dsn := fmt.Sprintf("%s://%s", dbDriver, dbSource)

	dbc, err := sql.Open(dbDriver, dsn)

	if err != nil {
		log.Fatalln("failed when connecting to database " + err.Error())
	}

	if err = dbc.Ping(); err != nil {
		log.Fatalln("failed when ping to database " + err.Error())
	}
	return dbc

}
