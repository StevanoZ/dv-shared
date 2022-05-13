package shrd_utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	ID uuid.UUID
}

func TestConvertInterface(t *testing.T) {
	copy := testData{}

	data := testData{
		ID: uuid.New(),
	}

	t.Run("Success converting", func(t *testing.T) {
		err := ConvertInterface(data, &copy)

		assert.NoError(t, err)
		assert.Equal(t, data.ID, copy.ID)
	})

	t.Run("Failed when converting", func(t *testing.T) {
		err := ConvertInterface(data, &[]testData{})

		assert.Error(t, err)
	})

	t.Run("Failed when unmarshal", func(t *testing.T) {
		err := ConvertInterface(func() {}, &testData{})

		assert.Error(t, err)
	})
}
