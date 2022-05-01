package shrd_utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ParamType interface {
	uuid.UUID | string
}

func ValidateUrlParamUUID(r *http.Request, paramName string) uuid.UUID {
	param := chi.URLParam(r, paramName)

	uuid, err := uuid.Parse(param)

	if err != nil {
		PanicIfError(CustomErrorWithTrace(err, fmt.Sprintf("invalid %s param", paramName), 400))
	}

	return uuid
}

func ValidateQueryParamInt(r *http.Request, queryName string) int {
	query := r.URL.Query().Get(queryName)

	queryInt, err := strconv.Atoi(query)

	if err != nil {
		PanicIfError(CustomErrorWithTrace(err, fmt.Sprintf("invalid %s query", queryName), 400))
	}

	if queryInt < 0 {
		PanicIfError(CustomError(fmt.Sprintf("invalid %s query", queryName), 400))
	}
	
	return queryInt
}

func CheckIsAuthorize(r *http.Request, accessId uuid.UUID) {
	tokenPayload := GetRequestCtx(r, TOKEN_PAYLOAD)

	if tokenPayload.UserId != accessId {
		PanicIfError(CustomError("not authorize to perform this operation", 403))
	}
}
