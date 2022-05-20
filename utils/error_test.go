package shrd_utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errApp = errors.New("panic if app error")
)

func TestPanicIfError(t *testing.T) {
	assert.PanicsWithValue(t, AppError{
		Message:    "panic if error",
		StatusCode: 500,
	}, func() {
		err := errors.New("panic if error")
		PanicIfError(err)
	})
}

func TestPanicIfAppError(t *testing.T) {
	assert.PanicsWithValue(t, AppError{
		Message:    fmt.Sprintf("%s|testing", errApp),
		StatusCode: 422,
	}, func() {
		err := errApp
		PanicIfAppError(err, "testing", 422)
	})
}

func TestPanicAppError(t *testing.T) {
	assert.PanicsWithValue(t, AppError{
		Message:    fmt.Sprintf("|%s", errApp.Error()),
		StatusCode: 422,
	}, func() {
		PanicAppError(errApp.Error(), 422)
	})
}

func TestPanicValidationError(t *testing.T) {
	defer func() {
		err := recover()

		validationErrors, isValidationErrors := err.(ValidationErrors)
		assert.Equal(t, true, isValidationErrors)
		assert.Equal(t, validationErrors.StatusCode, 400)
		assert.Equal(t, "validation errors", validationErrors.Errors[0].Message)
	}()

	validationError := []ValidationError{{Message: "validation errors"}}
	PanicValidationError(validationError, 400)
}

func TestValidateStruct(t *testing.T) {
	defer func() {
		err := recover()

		validationErrors, isValidationErrors := err.(ValidationErrors)
		assert.Equal(t, true, isValidationErrors)
		assert.Equal(t, 400, validationErrors.StatusCode)
		assert.Equal(t, "Field validation for 'Email' failed on the 'email' tag", validationErrors.Errors[0].Message)

	}()
	type testInput struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"min=3,max=10"`
		Password string `json:"password" validate:"min=8,max=15"`
	}

	input := testInput{
		Email:    "test@test",
		Username: "Testing",
		Password: "xxxxxxxx",
	}

	ValidateStruct(&input)
}

func TestValidateBodyPayload(t *testing.T) {
	type testInput struct {
		Success bool `json:"success"`
	}
	input := testInput{
		Success: true,
	}
	body, err := json.Marshal(input)
	assert.NoError(t, err)
	reader := bytes.NewReader(body)
	var output testInput

	ValidateBodyPayload(io.NopCloser(reader), &output)
	assert.Equal(t, true, output.Success)
}

func TestDeferCheck(t *testing.T) {
	DeferCheck(func() error {
		return errors.New("error")
	})
}
