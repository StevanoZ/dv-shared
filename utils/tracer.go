package shrd_utils

import (
	"context"
	"errors"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func SetupTracer(ctx context.Context, identifier string) func() {
	trac, errTrac := tracer.StartSpanFromContext(ctx, identifier)
	LogIfError(errTrac.Err())

	return func() {
		err := recover()

		if err != nil {
			appErr, isAppErr := err.(AppError)
			unknownErr, isUnknownErr := err.(error)
			if isAppErr {
				trac.Finish(tracer.WithError(errors.New(appErr.Message)))
			}

			if isUnknownErr {
				trac.Finish(tracer.WithError(unknownErr))
			}

			panic(err)
		} else {
			trac.Finish(tracer.NoDebugStack())
		}
	}
}
