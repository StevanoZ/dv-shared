package shrd_utils

import (
	"context"
	"errors"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func createTracer(ctx context.Context, identifier string) (context.Context, ddtrace.Span) {
	trac, tracErr := tracer.StartSpanFromContext(ctx, identifier)
	LogIfError(tracErr.Err())

	return tracer.ContextWithSpan(ctx, trac), trac
}

func CreateAndCheckTracerSvc(ctx context.Context, identifier string) (context.Context, func()) {
	svcCtx, trac := createTracer(ctx, identifier)

	return svcCtx, func() {
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
func CreateAndCheckTracer(ctx context.Context, identifier string) (context.Context, func(error)) {
	trac, tracErr := tracer.StartSpanFromContext(ctx, identifier)
	LogIfError(tracErr.Err())

	return tracer.ContextWithSpan(ctx, trac), func(err error) {
		checkTracer(trac, err)
	}
}

func checkTracer(trac ddtrace.Span, err error) {
	if err != nil {
		go trac.Finish(tracer.WithError(err))
	} else {
		go trac.Finish(tracer.NoDebugStack())
	}

}

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
