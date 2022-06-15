package shrd_utils

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupTracer(t *testing.T) {
	ctx := context.Background()

	t.Run("Should send error trace (app error)", func(t *testing.T) {
		assert.Panics(t, func() {
			defer SetupTracer(ctx, "testing-tracer")()

			PanicAppError("failed", 400)
		})
	})
	t.Run("Should send error trace (unknown error)", func(t *testing.T) {
		assert.Panics(t, func() {
			defer SetupTracer(ctx, "testing-tracer")()

			panic(errors.New("unknown error"))
		})
	})

	t.Run("Should send default trace", func(t *testing.T) {
		assert.NotPanics(t, func() {
			defer SetupTracer(ctx, "testing-tracer")()
		})
	})
}
