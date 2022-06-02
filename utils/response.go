package shrd_utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// FOR TESTING PURPOSE
type ResponseMap struct {
	Success    bool                   `json:"success"`
	StatusCode int                    `json:"statusCode"`
	Data       map[string]interface{} `json:"data"`
}

// FOR TESTING PURPOSE
type ResponseSlice struct {
	Success    bool                     `json:"success"`
	StatusCode int                      `json:"statusCode"`
	Data       []map[string]interface{} `json:"data"`
}

func GenerateSuccessResp(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success:    true,
		StatusCode: statusCode,
		Data:       data,
	}

	responseEncode, err := json.Marshal(response)
	PanicIfAppError(err, "failed when marshal response", 500)

	_, err = w.Write(responseEncode)
	PanicIfAppError(err, "failed when write success response", 500)
}

func GenerateErrorResp(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success:    false,
		StatusCode: statusCode,
		Data:       data,
	}

	responseEncode, err := json.Marshal(response)
	PanicIfAppError(err, "failed when marshar response", 500)

	_, err = w.Write(responseEncode)
	PanicIfAppError(err, "failed when write success response", 500)
}
