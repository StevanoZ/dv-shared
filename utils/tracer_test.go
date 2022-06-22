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

func TestCreateAndCheckTracer(t *testing.T) {
	ctx := context.Background()
	ctx, check := CreateAndCheckTracer(ctx, "testing")
	assert.NotNil(t, ctx)
	check(errors.New("error"))
	check(nil)
}

func TestCreateAndCheckTracerSvc(t *testing.T) {
	ctx := context.Background()

	t.Run("No error", func(t *testing.T) {
		ctx, check := CreateAndCheckTracerSvc(ctx, "testing")
		assert.NotNil(t, ctx)
		defer check()
	})
	t.Run("Should send error trace (app error)", func(t *testing.T) {
		assert.Panics(t, func() {
			ctx, check := CreateAndCheckTracerSvc(ctx, "testing")
			defer check()
			assert.NotNil(t, ctx)

			PanicAppError("error", 400)

		})
	})
	t.Run("Should send error trace (unknown error)", func(t *testing.T) {
		assert.Panics(t, func() {
			ctx, check := CreateAndCheckTracerSvc(ctx, "testing")
			assert.NotNil(t, ctx)
			defer check()

			panic(errors.New("unknown error"))
		})
	})
}
