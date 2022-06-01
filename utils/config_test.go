package shrd_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBaseConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		config := LoadBaseConfig("../app", "test")
		assert.NotNil(t, config)
	})
}
