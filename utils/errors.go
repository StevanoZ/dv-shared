package shrd_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AppError struct {
	Message    string
	StatusCode int
}

func (ae *AppError) Error() string {
	return fmt.Sprintf("app error: status code %d, message %s", ae.StatusCode, ae.Message)
}

type ValidationError struct {
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error: message %s", ve.Message)
}

type ValidationErrors struct {
	Errors     []ValidationError
	StatusCode int
}

func (ve *ValidationErrors) Error() string {
	return fmt.Sprintf("validation errors: status code %d, message %s", ve.StatusCode, ve.Errors[0].Message)
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
	err := json.NewDecoder(body).Decode(output)
	PanicIfAppError(err, "failed when decode body payload", 400)

	ValidateStruct(output)
}

func DeferCheck(function func() error) {
	if err := function(); err != nil {
		LogError("defer error", zap.Error(err))
	}
}

func LogIfError(err error) {
	if err != nil {
		LogError("error occurred", zap.Error(err))
	}
}

func LogAndPanicIfError(err error, message string) {
	if err != nil {
		errMsg := fmt.Sprintf("%s :%v", message, err)
		LogError(errMsg, zap.Error(err))
		panic(err)
	}
}
