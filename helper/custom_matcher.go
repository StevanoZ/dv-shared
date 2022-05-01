package shrd_helper

import (
	"context"
	"fmt"

	shrd_token "github.com/StevanoZ/dv-shared/token"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

type tokenPayloadContext struct {
	userId uuid.UUID
}

func TokenPayloadContextMatcher(userId uuid.UUID) gomock.Matcher {
	return tokenPayloadContext{userId: userId}
}

func (e tokenPayloadContext) Matches(x interface{}) bool {
	ctx, isOk := x.(context.Context)
	if !isOk {
		return false
	}

	payload := ctx.Value(shrd_utils.TOKEN_PAYLOAD).(*shrd_token.Payload)

	return e.userId == payload.UserId
}

func (e tokenPayloadContext) String() string {
	return fmt.Sprintf("matches context with userId: %s", e.userId)
}
