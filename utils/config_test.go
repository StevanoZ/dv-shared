package shrd_utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBaseConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		config := LoadBaseConfig("../app", "test")
		assert.NotNil(t, config)
	})
}

func TestCheckAndSetConfig(t *testing.T) {
	t.Run("Load local config", func(t *testing.T) {
		assert.NotPanics(t, func() {
			config := CheckAndSetConfig("../app", "app")
			assert.NotNil(t, config)
			assert.Equal(t, "local", config.Environment)
			assert.Equal(t, "host.docker.internal:8085", config.PS_PUBSUB_EMULATOR_HOST)
		})
	})

	t.Run("Load test config", func(t *testing.T) {
		os.Setenv("ENVIRONMENT", "test")
		assert.NotPanics(t, func() {
			config := CheckAndSetConfig("../app", "app")
			assert.NotNil(t, config)
			assert.Equal(t, "test", config.Environment)
			assert.Equal(t, "localhost:8085", config.PS_PUBSUB_EMULATOR_HOST)
		})
	})
}
