package shrd_utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(dbDriver, dbSource string) *sql.DB {
	dsn := fmt.Sprintf("%s://%s", dbDriver, dbSource)

	dbc, err := sql.Open(dbDriver, dsn)
	LogAndPanicIfError(err, "failed when connecting to database")

	err = dbc.Ping()
	LogAndPanicIfError(err, "failed when ping to database")

	return dbc
}
