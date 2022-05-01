package shrd_middleware

import (
	"net/http"

	"strings"

	shrd_token "github.com/StevanoZ/dv-shared/token"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

type AuthMiddleware interface {
	CheckIsAuthenticated(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc
}

type AuthMiddlewareImpl struct {
	tokenMaker shrd_token.Maker
}

func NewAuthMiddleware(tokenMaker shrd_token.Maker) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		tokenMaker: tokenMaker,
	}
}

func (m *AuthMiddlewareImpl) CheckIsAuthenticated(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if header == "" || !strings.Contains(header, "Bearer ") {
			shrd_utils.PanicIfError(shrd_utils.CustomError("invalid token", 401))
		}
		token := strings.Split(header, " ")[1]

		payload, err := m.tokenMaker.VerifyToken(token)

		if err != nil {
			shrd_utils.PanicIfError(shrd_utils.CustomErrorWithTrace(err, "invalid token", 401))
		}

		if payload.Status != "active" {
			shrd_utils.PanicIfError(shrd_utils.CustomError("inactive user can't access this route", 403))
		}

		ctx := shrd_utils.AppendRequestCtx(r, shrd_utils.TOKEN_PAYLOAD, payload)
		handler(w, r.WithContext(ctx))
	}
}
