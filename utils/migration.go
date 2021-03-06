package shrd_utils

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration(db *sql.DB, config *BaseConfig) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	LogAndPanicIfError(err, "cannot create new migrate instance")

	m, err := migrate.NewWithDatabaseInstance(
		config.MIGRATION_URL,
		config.DBDriver, driver)
	LogAndPanicIfError(err, "failed to run migrate up")

	m.Up()
}
