package shrd_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Message    string
	StatusCode int
}

type ValidationError struct {
	Message string
}

type ValidationErrors struct {
	Errors     []ValidationError
	StatusCode int
}

func CustomError(message string, statusCode int) error {
	return fmt.Errorf("|%s<->%d", message, statusCode)
}

func CustomErrorWithTrace(err error, message string, statusCode int) error {
	return fmt.Errorf("%s|%s<->%d", err.Error(), message, statusCode)
}

func PanicIfError(err error) {
	if err != nil {
		customError := strings.Split(err.Error(), "<->")
		message := customError[0]
		statusCode := 500

		if len(customError) > 1 {
			statusCode, _ = strconv.Atoi(customError[1])
		}

		appErr := AppError{
			Message:    message,
			StatusCode: statusCode,
		}
		panic(appErr)
	}
}

func PanicIfAppError(err error, message string, statusCode int) {
	if err != nil {
		customErr := CustomErrorWithTrace(err, message, statusCode)
		PanicIfError(customErr)
	}
}

func PanicAppError(message string, statusCode int) {
	customErr := CustomError(message, statusCode)
	PanicIfError(customErr)
}

func PanicValidationError(errors []ValidationError, statusCode int) {
	validationErrors := ValidationErrors{
		Errors:     errors,
		StatusCode: statusCode,
	}
	panic(validationErrors)
}

func ValidateStruct(data interface{}) {
	var validationErrors []ValidationError
	validate := validator.New()
	errorValidate := validate.Struct(data)

	if errorValidate != nil {
		for _, err := range errorValidate.(validator.ValidationErrors) {
			var validationError ValidationError
			validationError.Message = strings.Split(err.Error(), "Error:")[1]
			validationErrors = append(validationErrors, validationError)
		}
		PanicValidationError(validationErrors, 400)
	}
}

func ValidateBodyPayload(body io.ReadCloser, output interface{}) {
	json.NewDecoder(body).Decode(output)
	ValidateStruct(output)
}

func DeferCheck(function func() error) {
	if err := function(); err != nil {
		log.Println("defer error:", err)
	}
}

func LogIfError(err error) {
	if err != nil {
		log.Println("error occured: ", err)
	}
}
