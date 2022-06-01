package shrd_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunMigration(t *testing.T) {
	config := LoadBaseConfig("../app", "test")
	db := ConnectDB(config.DBDriver, config.DBSource)

	assert.NotPanics(t, func() {
		config.MIGRATION_URL = "file://../db/migration"
		RunMigration(db, config)
	})
}
