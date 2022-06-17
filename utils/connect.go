package shrd_utils

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

func ConnectDB(dbDriver, dbSource string) *sql.DB {
	dsn := fmt.Sprintf("%s://%s", dbDriver, dbSource)

	sqltrace.Register("postgres", &pq.Driver{})
	dbc, err := sql.Open(dbDriver, dsn)
	LogAndPanicIfError(err, "failed when connecting to database")

	err = dbc.Ping()
	LogAndPanicIfError(err, "failed when ping to database")

	return dbc
}
