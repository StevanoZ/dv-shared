package shrd_utils

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func SetUpTracerSpan(ctx context.Context, identifier string) ddtrace.Span {
	trac, errTrac := tracer.StartSpanFromContext(ctx, identifier)
	LogIfError(errTrac.Err())

	return trac
}
