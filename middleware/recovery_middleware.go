package shrd_middleware

import (
	"log"
	"net/http"
	"strings"

	shrd_utils "github.com/StevanoZ/dv-shared/utils"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				var errorMsgs []map[string]interface{}
				var statusCode int

				appErr, isAppErr := err.(shrd_utils.AppError)
				validationErr, isValidationErr := err.(shrd_utils.ValidationErrors)

				if isAppErr {
					messages := strings.Split(appErr.Message, "|")
					log.Println("APP ERROR (PANIC)", messages[0])
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
				} else {
					log.Println("UNKNOWN ERROR (PANIC)", err)
					errorMsgs = []map[string]interface{}{
						{"message": "internal server error"},
					}
					statusCode = 500
				}
				shrd_utils.GenerateErrorResp(w, errorMsgs, statusCode)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
