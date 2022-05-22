package shrd_helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	shrd_token "github.com/StevanoZ/dv-shared/token"
	shrd_utils "github.com/StevanoZ/dv-shared/utils"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type TestCaseHandler struct {
	Name          string
	SetHeaders    func(req *http.Request)
	Payload       interface{}
	Method        string
	ReqUrl        string
	BuildStub     func(input interface{}, stubs ...interface{})
	CheckResponse func(recorder *httptest.ResponseRecorder, expected interface{})
}

func SetupRequest(t *testing.T, r *chi.Mux, tc TestCaseHandler, stubs ...interface{}) {
	var req *http.Request
	var err error
	recorder := httptest.NewRecorder()
	input := tc.Payload

	if tc.BuildStub != nil {
		tc.BuildStub(input, stubs...)
	}

	formData, isFormData := input.(*bytes.Buffer)
	if isFormData {
		req, err = http.NewRequest(tc.Method, tc.ReqUrl, formData)
		assert.NoError(t, err)
	} else {
		body, err := json.Marshal(input)
		assert.NoError(t, err)
		req, err = http.NewRequest(tc.Method, tc.ReqUrl, bytes.NewReader(body))
		assert.NoError(t, err)
	}

	if tc.SetHeaders != nil {
		tc.SetHeaders(req)
	}

	r.ServeHTTP(recorder, req)
	tc.CheckResponse(recorder, input)
}

func SetHeaderApplicationJson(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
}

func SetHeaderMultiPartForm(req *http.Request, contentType string) {
	req.Header.Add("Content-Type", contentType)
}

func SetAuthorizationHeader(req *http.Request, symmetricKey string, userId uuid.UUID) {
	tokenMaker, _ := shrd_token.NewPasetoMaker(&shrd_utils.BaseConfig{TokenSymmetricKey: symmetricKey})

	token, _, _ := tokenMaker.CreateToken(shrd_token.PayloadParams{
		UserId: userId,
		Email:  shrd_utils.RandomEmail(),
		Status: "active",
	}, time.Hour)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
}

func ParseResponseBody(body *bytes.Buffer, output interface{}) {
	json.NewDecoder(body).Decode(output)
}

func ParseInterfaceToMap(data interface{}, field string) map[string]interface{} {
	return data.(map[string]interface{})[field].(map[string]interface{})
}

func ParseInterfaceToSlice(data interface{}, field string) []interface{} {
	return data.(map[string]interface{})[field].([]interface{})
}

func ParseInterfaceToString(data interface{}, field string) string {
	return data.(map[string]interface{})[field].(string)
}

func ParseErrorMessage(data interface{}) string {
	errors := data.([]interface{})

	errorsResp := ""
	for i, err := range errors {
		errMap := err.(map[string]interface{})
		if i > 0 {
			errorsResp = fmt.Sprintf("%s, %s", errorsResp, errMap["message"])

		} else {
			errorsResp = errMap["message"].(string)
		}

	}
	return errorsResp
}

func CreateFormFile(n int, filename string) (*bytes.Buffer, string) {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	for i := 0; i < n; i++ {
		newFileName := fmt.Sprintf("%s-%d", filename, i+1)
		formFile, _ := writer.CreateFormFile("files", newFileName)
		formFile.Write([]byte(fmt.Sprintf("JUST FOR TESTING %s", newFileName)))
	}
	writer.Close()

	return form, writer.FormDataContentType()
}

func CreateFilesHeader(n int, filename string) []*multipart.FileHeader {
	formFile, header := CreateFormFile(n, filename)

	req := httptest.NewRequest(http.MethodPost, "/upload", formFile)
	defer req.Body.Close()
	req.Header.Add("Content-Type", header)

	req.ParseMultipartForm(5242880)
	filesHeader := req.MultipartForm.File["files"]

	return filesHeader
}

func CheckTokenPayloadCtx(ctx context.Context) gomock.Matcher {

	return gomock.AssignableToTypeOf(context.WithValue(ctx, gomock.Nil(), ""))
}

func CheckResponse200(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, true, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 200, responseMap.StatusCode)
		assert.Equal(t, true, responseMap.Success)
	}
}

func CheckResponse201(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 201, response.StatusCode)
		assert.Equal(t, true, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 201, responseMap.StatusCode)
		assert.Equal(t, true, responseMap.Success)
	}
}

func CheckResponse400(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, false, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 400, responseMap.StatusCode)
		assert.Equal(t, false, responseMap.Success)
	}
}

func CheckResponse401(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 401, response.StatusCode)
		assert.Equal(t, false, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 401, responseMap.StatusCode)
		assert.Equal(t, false, responseMap.Success)
	}
}

func CheckResponse404(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 404, response.StatusCode)
		assert.Equal(t, false, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 404, responseMap.StatusCode)
		assert.Equal(t, false, responseMap.Success)
	}
}

func CheckResponse422(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 422, response.StatusCode)
		assert.Equal(t, false, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 422, responseMap.StatusCode)
		assert.Equal(t, false, responseMap.Success)
	}
}

func CheckResponse500(t *testing.T, apiResponse interface{}) {
	// JUST HAVE 2 CONDITIONS --> Response or ResponseMap (for testing purpose)
	response, isResponse := apiResponse.(shrd_utils.Response)
	responseMap, isResponseMap := apiResponse.(shrd_utils.ResponseMap)

	if isResponse {
		assert.Equal(t, 500, response.StatusCode)
		assert.Equal(t, false, response.Success)
	}

	if isResponseMap {
		assert.Equal(t, 500, responseMap.StatusCode)
		assert.Equal(t, false, responseMap.Success)
	}
}
