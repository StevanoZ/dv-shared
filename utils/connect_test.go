package shrd_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	config := LoadBaseConfig("../app", "test")
	assert.NotPanics(t, func() {
		db := ConnectDB(config.DBDriver, config.DBSource)
		assert.NotNil(t, db)
	})
}
