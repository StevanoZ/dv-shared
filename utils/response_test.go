package shrd_utils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
func TestGenerateSuccessResp(t *testing.T) {
	recorder := setUpRecorder()
	data := "Test Success"
	GenerateSuccessResp(recorder, data, 201)

	var resp Response
	json.NewDecoder(recorder.Body).Decode(&resp)

	assert.Equal(t, true, resp.Success)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 201, recorder.Result().StatusCode)
	assert.Equal(t, "Test Success", resp.Data)
}

func TestGenerateErrorResp(t *testing.T) {
	recorder := setUpRecorder()

	data := "Test Failed"
	GenerateErrorResp(recorder, data, 400)

	var resp Response
	json.NewDecoder(recorder.Body).Decode(&resp)

	assert.Equal(t, false, resp.Success)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, 400, recorder.Result().StatusCode)
	assert.Equal(t, "Test Failed", resp.Data)
}
