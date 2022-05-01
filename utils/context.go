package shrd_utils

import (
	"context"
	"net/http"

	shrd_token "github.com/StevanoZ/dv-shared/token"
)

type key string

const TOKEN_PAYLOAD key = "token-payload"

func AppendRequestCtx(r *http.Request, ctxKey key, input interface{}) context.Context {
	return context.WithValue(r.Context(), ctxKey, input)
}

func GetRequestCtx(r *http.Request, ctxKey key) *shrd_token.Payload {
	return r.Context().Value(ctxKey).(*shrd_token.Payload)
}
