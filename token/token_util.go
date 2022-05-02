package shrd_token

import (
	"context"
	"net/http"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/google/uuid"
)

type key string

const TOKEN_PAYLOAD key = "token-payload"

func AppendRequestCtx(r *http.Request, ctxKey key, input interface{}) context.Context {
	return context.WithValue(r.Context(), ctxKey, input)
}

func GetRequestCtx(r *http.Request, ctxKey key) *Payload {
	return r.Context().Value(ctxKey).(*Payload)
}

func CheckIsAuthorize(r *http.Request, accessId uuid.UUID) {
	tokenPayload := GetRequestCtx(r, TOKEN_PAYLOAD)

	if tokenPayload.UserId != accessId {
		shrd_utils.PanicIfError(shrd_utils.CustomError("not authorize to perform this operation", 403))
	}
}
