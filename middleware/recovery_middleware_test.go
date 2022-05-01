package shrd_middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	shrd_helper "github.com/StevanoZ/dv-shared/helper"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"

	"github.com/stretchr/testify/assert"
)

const (
	APP_ERROR        = "app error"
	VALIDATION_ERROR = "invalid username, invalid email"
	UNKNOWN_ERROR    = "internal server error"
)

type TestCaseRecoveryMiddleware struct {
	Name          string
	Handler       http.HandlerFunc
	CheckResponse func(recorder *httptest.ResponseRecorder)
}

func setUpRecorder() (*httptest.ResponseRecorder, *http.Request) {
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	recorder := httptest.NewRecorder()

	return recorder, req
}

func TestRecoveryMiddleware(t *testing.T) {
	testCasesRecovery := []TestCaseRecoveryMiddleware{
		{
			Name: "App Error",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				shrd_utils.PanicIfError(shrd_utils.CustomError(APP_ERROR, 422))
			}),
			CheckResponse: func(recorder *httptest.ResponseRecorder) {
				resp := shrd_utils.Response{}

				shrd_helper.ParseResponseBody(recorder.Body, &resp)

				shrd_helper.CheckResponse422(t, resp)

				appErr := shrd_helper.ParseErrorMessage(resp.Data)
				assert.Equal(t, APP_ERROR, appErr)
			},
		},
		{
			Name: "Validation Error",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				validationErrors := []shrd_utils.ValidationError{
					{
						Message: "invalid username",
					},
					{
						Message: "invalid email",
					},
				}
				shrd_utils.PanicValidationError(validationErrors, 400)
			},
			CheckResponse: func(recorder *httptest.ResponseRecorder) {
				resp := shrd_utils.Response{}

				shrd_helper.ParseResponseBody(recorder.Body, &resp)

				shrd_helper.CheckResponse400(t, resp)

				validationErr := shrd_helper.ParseErrorMessage(resp.Data)
				assert.Equal(t, "invalid username, invalid email", validationErr)
			},
		},
		{
			Name: "Unknown Error",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				panic("should not sent this error to the client")
			},
			CheckResponse: func(recorder *httptest.ResponseRecorder) {
				resp := shrd_utils.Response{}

				shrd_helper.ParseResponseBody(recorder.Body, &resp)

				shrd_helper.CheckResponse500(t, resp)

				unknownErr := shrd_helper.ParseErrorMessage(resp.Data)
				assert.Equal(t, UNKNOWN_ERROR, unknownErr)
			},
		},
	}

	recorder, req := setUpRecorder()

	for i := range testCasesRecovery {
		tc := testCasesRecovery[i]

		handler := Recovery(tc.Handler)
		handler.ServeHTTP(recorder, req)
		tc.CheckResponse(recorder)
	}

}
