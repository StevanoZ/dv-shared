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
				go trac.Finish(tracer.WithError(errors.New(appErr.Message)))
			}

			if isUnknownErr {
				go trac.Finish(tracer.WithError(unknownErr))
			}

			panic(err)
		} else {
			go trac.Finish(tracer.NoDebugStack())
		}
	}
}
