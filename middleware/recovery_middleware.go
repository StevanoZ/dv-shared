package shrd_middleware

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {	
				var errorMsgs []map[string]interface{}
				var statusCode int
				var traceName string
				var errTrace string

				appErr, isAppErr := err.(shrd_utils.AppError)
				validationErr, isValidationErr := err.(shrd_utils.ValidationErrors)

				if isAppErr {
					messages := strings.Split(appErr.Message, "|")
					log.Println("APP ERROR (PANIC)", messages[0])

					traceName = "app error"
					errTrace = messages[0]
					errorMsgs = []map[string]interface{}{
						{"message": messages[1]},
					}
					statusCode = appErr.StatusCode
				} else if isValidationErr {
					log.Println("VALIDATION ERROR (PANIC)", validationErr)

					for _, err := range validationErr.Errors {
						errorMsg := map[string]interface{}{
							"message": err.Message,
						}
						errorMsgs = append(errorMsgs, errorMsg)
					}
					statusCode = validationErr.StatusCode

					traceName = "validation error"
					errStr, _ := json.Marshal(errorMsgs)
					errTrace = string(errStr)
				} else {
					log.Println("UNKNOWN ERROR (PANIC)", err)
					errorMsgs = []map[string]interface{}{
						{"message": "internal server error"},
					}
					statusCode = 500

					traceName = "unknown error"
					errPanic, isErrPanic := err.(error)
					errTrace = "internal server error"

					if isErrPanic {
						errTrace = errPanic.Error()
					}
				}

				trc := shrd_utils.SetUpTracerSpan(r.Context(), traceName)
				trc.Finish(
					tracer.WithError(errors.New(errTrace)),
				)

				shrd_utils.GenerateErrorResp(w, errorMsgs, statusCode)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
